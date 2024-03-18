defmodule Connections.RegistryCleaner do
  use GenServer

  require Logger

  def start_link(_) do
    GenServer.start_link(__MODULE__, %{})
  end

  @impl true
  def init(state) do
    :timer.send_interval(1_000, :work)
    {:ok, state}
  end

  @impl true
  def handle_info(:work, state) do
    do_recurrent_thing(state)
    {:noreply, state}
  end

  defp do_recurrent_thing(_state) do
    Connections.Registry.get_all
    |> Enum.map(fn {k, m} ->
      if not m.connected do
        td = DateTime.diff(DateTime.now!("Etc/UTC"), m.last_seen)
        if td >= 5 do
          Connections.Registry.remove_entry(k)
        end
      end
    end)
  end
end
