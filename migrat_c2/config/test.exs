import Config

# We don't run a server during test. If one is required,
# you can enable the server option below.
config :migrat_c2, MigratC2Web.Endpoint,
  http: [ip: {127, 0, 0, 1}, port: 4002],
  secret_key_base: "5ZY0RtLsvRtlKMLPg1MKKq+doItgbyk2lxV0cmCCtonVIzrkIgutcDsvsG0Z2TRM",
  server: false

# In test we don't send emails.
config :migrat_c2, MigratC2.Mailer,
  adapter: Swoosh.Adapters.Test

# Print only warnings and errors during test
config :logger, level: :warn

# Initialize plugs at runtime for faster test compilation
config :phoenix, :plug_init_mode, :runtime
