defmodule Connections.Registry do
  def start do
    :ets.new(:connections_registry, [:named_table, :public])
  end

  def get_all do
    :ets.tab2list(:connections_registry)
  end

  def get(id) do
    case :ets.lookup(:connections_registry, id) do
      [{_, m}] -> {:ok, m}
      [] -> {:err, []}
    end
  end

  def update_running_modules(id, modules) do
    case :ets.lookup(:connections_registry, id) do
      [{_, m}] -> :ets.insert(:connections_registry, {id, %{m | modules: modules}})
      [] -> Logger.error("Connection #{id} not found")
    end
  end

  def insert_shell_line(id, line) do
    [{k, map}] = :ets.lookup(:connections_registry, id)
    map = %{ map | shell_output: [line | map.shell_output]}
    :ets.insert(:connections_registry, {k, map})
    MigratC2.LiveUpdate.notify_update()
  end

  def new_heartbeat(ident) do
    case :ets.lookup(:connections_registry, ident.id) do
      [] ->
        map = %{
          ident: ident,
          connected: false,
          wants_connection: false,
          last_seen: DateTime.now!("Etc/UTC"),
          handler: nil,
          shell_output: [],
          modules: [],
        }
        :ets.insert(:connections_registry, {ident.id, map})
        MigratC2.LiveUpdate.notify_update()
        map
      [{k, m}] ->
        m = %{ m | last_seen: DateTime.now!("Etc/UTC") }
        :ets.insert(:connections_registry, {k, m})
        m
    end
  end

  def connection_established(ident) do
    [{k, map}] = :ets.lookup(:connections_registry, ident.id)
    map = %{ map | connected: true, handler: self(), wants_connection: false}
    :ets.insert(:connections_registry, {k, map})
    MigratC2.LiveUpdate.notify_update()
  end

  def connection_closed(id) do
    [{k, map}] = :ets.lookup(:connections_registry, id)
    map = %{ map | connected: false, handler: nil, wants_connection: false, last_seen: DateTime.now!("Etc/UTC")}
    :ets.insert(:connections_registry, {k, map})
    MigratC2.LiveUpdate.notify_update()
  end

  def request_connection(id) do
    [{k, map}] = :ets.lookup(:connections_registry, id)
    map = %{ map | wants_connection: true}
    :ets.insert(:connections_registry, {k, map})
  end

  def remove_entry(id) do
   :ets.delete(:connections_registry, id)
   MigratC2.LiveUpdate.notify_update()
  end

end
