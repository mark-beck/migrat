defmodule MigratC2.LiveUpdate do
  require Logger

  def subscibe_new_connection do
    Phoenix.PubSub.subscribe(MigratC2.PubSub, "connections")
  end

  def subscribe_id(id) do
    Phoenix.PubSub.subscribe(MigratC2.PubSub, "connections:#{id}")
  end

  def subscribe_id_cmd(id) do
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

  def module_output(id, name, data) do
    Logger.info("command_output notify on id #{id}")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections:#{id}", {:module_output, name, data})
  end

  def module_list(id, data) do
    Logger.info("module_list notify on id #{id}")
    Phoenix.PubSub.broadcast(MigratC2.PubSub, "connections:#{id}", {:module_list, data})
  end


end
