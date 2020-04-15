require 'socket'
require 'minitest/autorun'

def sh(command)
  result = `#{command}`
  raise "#{command} failed\n#{result}" unless $?.success?

  result
end

describe "metrics" do
  it "sends to datadog" do
    sh("go build .")
    u1 = UDPSocket.new
    port = 1234
    u1.bind("localhost", port)
    sh("echo 'this is bad' | STATSD_HOST='localhost' STATSD_PORT=#{port} ./logs_to_datadog_metrics")
    u1.recvfrom(1000).first.must_equal "foo.bad:1|c|#tag:bad"
  end
end
