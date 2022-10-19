defmodule MigratC2Web.ConnectionsLive.Endpoint.TerminalContainer do
  @moduledoc false

  use Phoenix.LiveComponent
  alias Phoenix.LiveView.JS

  import MigratC2Web.ConnectionsLive.Endpoint, only: [terminal: 1]

  def render(%{id: id, connection: connection} = assigns) do

    box_id = "#{id}_box"

    ~H"""
    <div id={box_id} style="border: 2px solid black">
      <div class="terminal_container" style="position: relative; border: 2px solid lightgray;">
        <.terminal id={@id} connection={connection} />
      </div>
    </div>
    """
  end
end
