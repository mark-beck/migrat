defmodule MigratC2.Inputs.ShellCommand do
  alias MigratC2.Inputs.ShellCommand
  import Ecto.Changeset

  defstruct [:command]
  @types %{command: :string}

  def changeset(%ShellCommand{} = shell_command, attrs \\ %{}) do
    {shell_command, @types}
    |> cast(attrs, [:command])
  end

end
