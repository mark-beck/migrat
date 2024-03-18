defmodule MigratC2.LiveUpdate do
  require Logger

  def subscibe_new_connection do
    Phoenix.PubSub.subscribe(MigratC2.PubSub, "connections")
  end

  def subscribe_id(id) do
    Phoenix.PubSub.subscribe(MigratC2.PubSub, "connections:#{id}")
  end

  def unsubscribe_new_connection do
    Phoenix.PubSub.unsubscribe(MigratC2.PubSub, "connections")
  end

  def unsubscribe_id(id) do
    Phoenix.PubSub.unsubscribe(MigratC2.PubSub, "connections:#{id}")
  end

  def notify_update() do
    Logger.info("notifying")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections", :update)
  end

  def new_screenshot(id, img) do
    Logger.info("screenshot notify")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections:#{id}", {:screenshot, id, img})
  end

  def new_shellmessage(id, msg) do
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections:#{id}", {:shellmessage, id, msg})
  end

  def command_response(id, type, data) do
    Logger.info("command_response notify")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections:#{id}", {:command_response, id, type, data})
  end


end
