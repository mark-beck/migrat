defmodule MigratC2Web.PageController do
  use MigratC2Web, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
