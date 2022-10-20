defmodule MigratC2Web.ConnectionsLive.Endpoint.Terminal do
  @moduledoc false
  require Logger

  use Phoenix.LiveComponent

  def render(assigns) do
    ~H"""
    <div class="terminal_container" style="position: relative; border: 2px solid lightgray;">
      <div phx-hook="Terminal" id={"#{@id}_#{@name}"}>
        <div class="xtermjs_container" phx-update="ignore" id={"xtermjs-container-#{@id}_#{@name}"}></div>
      </div>
    </div>

    """
  end

  def handle_event("command", command, socket) do
    IO.puts(command)
    Connections.Handler.send_module_input(socket.assigns.connection.handler, socket.assigns.name, command)
    {:noreply, socket}
  end

  def update(%{id: id, data: data}, socket) do
    Logger.info("sending command_out to terminal")
    {:ok, push_event(socket, "print", %{data: data})}
  end

  def update(assigns, socket) do
    Logger.info("Terminal got id: #{assigns.id}, name #{assigns.name}")

    socket = socket
    |> assign(assigns)
    {:ok, socket}
  end

end
