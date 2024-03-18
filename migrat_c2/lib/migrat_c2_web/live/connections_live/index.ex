defmodule MigratC2Web.ConnectionsLive.Index do
  use MigratC2Web, :live_view
  require Logger

  @impl true
  def mount(_params, _session, socket) do
    MigratC2.LiveUpdate.subscibe_new_connection()
    {
      :ok,
      socket
      |> assign(:connections, Connections.Registry.get_all
      |> Enum.map(fn {_, v} ->
        v
      end))
    }
  end

  @impl true
  def handle_event("select", %{"id"=>id}, socket) do
    Logger.info("selecting #{inspect{id}}")
    socket = push_navigate socket, to: ~p"/#{id}"
    {:noreply, socket}
  end

  @impl true
  def handle_info(:update, socket) do
    socket = socket |> assign(:connections, Connections.Registry.get_all
    |> Enum.map(fn {_, v} ->
      v
    end))
    {:noreply, socket}
  end


end
