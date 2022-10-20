defmodule Connections.Translator do
  require Logger

  def read_message(socket, key) do
    with {:ok, header} <- :gen_tcp.recv(socket, 8),
         bytes_send <- :binary.decode_unsigned(header, :little),
         {:ok, data} <- :gen_tcp.recv(socket, bytes_send) do
      data = xor_message(data, key)

      message_type = Binary.last(data)
      Logger.info("got message type #{message_type}")

      data =
        data
        |> Binary.reverse()
        |> Binary.drop(1)
        |> Binary.reverse()

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
      0 -> {:ident, Migrat.Ident.decode(data)}
      1 -> {:heartbeat_response, Migrat.HeartbeatResponse.decode(data)}
      2 -> {:error, Migrat.Error.decode(data)}
      10 -> {:module_start, Migrat.ModuleStart.decode(data)}
      11 -> {:module_input, Migrat.ModuleInput.decode(data)}
      12 -> {:module_output, Migrat.ModuleOutput.decode(data)}
      13 -> {:module_list, Migrat.ModuleList.decode(data)}
      20 -> {:take_screenshot, Migrat.TakeScreenshot.decode(data)}
      21 -> {:screenshot, Migrat.Screenshot.decode(data)}
    end
  end

  defp encode_message(type, m) do
    case type do
      :ident -> Migrat.Ident.encode(m) <> <<0>>
      :heartbeat_response -> Migrat.HeartbeatResponse.encode(m) <> <<1>>
      :error -> Migrat.Error.encode(m) <> <<2>>
      :module_start -> Migrat.ModuleStart.encode(m) <> <<10>>
      :module_input -> Migrat.ModuleInput.encode(m) <> <<11>>
      :module_output -> Migrat.ModuleOutput.encode(m) <> <<12>>
      :module_list -> Migrat.ModuleList.encode(m) <> <<13>>
      :take_screenshot -> Migrat.TakeScreenshot.encode(m) <> <<20>>
      :screenshot -> Migrat.Screenshot.encode(m) <> <<21>>
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
