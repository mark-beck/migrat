defmodule MigratC2.LiveUpdate do
  require Logger

  def subscibe_new_connection do
    Phoenix.PubSub.subscribe(MigratC2.PubSub, "connections")
  end

  def subscribe_id(id) do
    Phoenix.PubSub.subscribe(MigratC2.PubSub, "connections:#{id}")
  end

  def unsubscripe_new_connection do
    Phoenix.PubSub.unsubscribe(MigratC2.PubSub, "connections")
  end

  def notify_update() do
    Logger.info("notifying")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections", :update)
  end

  def new_screenshot(id, img) do
    Logger.info("screenshot notify")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections:#{id}", {:screenshot, img})
  end


end
