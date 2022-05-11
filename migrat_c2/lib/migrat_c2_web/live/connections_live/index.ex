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
  def handle_info(:update, socket) do
    socket = socket |> assign(:connections, Connections.Registry.get_all
    |> Enum.map(fn {_, v} ->
      v
    end))
    {:noreply, socket}
  end

  @impl true
  def render(assigns) do
    ~H{
      <h1>Connections:</h1>
      <table>
      <%= for conn <- @connections do %>
      <tr>
        <td><%= conn.ident.computerName %></td>
        <td><%= conn.ident.campainId %></td>
        <td><%= live_redirect conn.ident.id, to: Routes.connections_endpoint_path(@socket, :index, conn.ident.id) %></td>
      </tr>
      <% end %>
      </table>
    }
  end

end
