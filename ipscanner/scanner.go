ctx := context.Background()

scanner := ipscanner.NewScanner(
	ipscanner.WithUseIPv4(false),
	ipscanner.WithUseIPv6(true),
	ipscanner.WithCidrList([]netip.Prefix{netip.ParsePrefix("2001:db8::/64")}),
	ipscanner.WithHttpPath("/some-path"),
	ipscanner.WithUserAgent("MyCustomUserAgent"),
	ipscanner.WithLogger(&slog.Logger{
		Handler: slog.NewJSONHandler(os.Stdout),
		Level:   slog.LevelInfo,
	}),
	ipscanner.WithInsecureSkipVerify(false),
	ipscanner.WithHostname("example.com"),
	ipscanner.WithPort(80),
	ipscanner.WithHTTPPing(),
	ipscanner.WithQUICPing(),
	ipscanner.WithTCPPing(),
	ipscanner.WithTLSPing(),
	ipscanner.WithConnectionTimeout(5*time.Second),
	ipscanner.WithHandshakeTimeout(3*time.Second),
	ipscanner.WithTlsVersion(tls.VersionTLS12),
)

scanner.Run(ctx)

availableIPs := scanner.GetAvailableIPs()
for _, ipInfo := range availableIPs {
	fmt.Println(ipInfo)
}
