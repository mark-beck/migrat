defmodule MigratC2.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Telemetry supervisor
      MigratC2Web.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: MigratC2.PubSub},
      # Start the Endpoint (http/https)
      MigratC2Web.Endpoint,
      # Start a worker by calling: MigratC2.Worker.start_link(arg)
      # {MigratC2.Worker, arg}

      {DynamicSupervisor, strategy: :one_for_one, name: Connections.HandlerSupervisor},
      {Task.Supervisor, name: Connections.ReadloopSupervisor},
      Supervisor.child_spec({Task, fn -> Connections.Server.accept(4040) end}, restart: :permanent),
      Connections.RegistryCleaner
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: MigratC2.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    MigratC2Web.Endpoint.config_change(changed, removed)
    :ok
  end
end
