require 'socket'
require 'minitest/autorun'

def sh(command)
  result = `#{command}`
  raise "#{command} failed\n#{result}" unless $?.success?

  result
end

sh("go build .")

describe "metrics" do
  def listen
    UDPSocket.open do |u1|
      u1.bind("localhost", 1234)
      ENV['STATSD_HOST'] = "localhost"
      ENV['STATSD_PORT'] = '1234'
      yield
      u1.recvfrom(1000).first
    end
  end

  it "sends to datadog" do
    listen { sh("echo 'this is bad' | ./logs_to_datadog_metrics") }.must_equal "foo.bad:1|c|#tag:bad"
  end

  it "does not need tags" do
    listen { sh("echo 'this is untagged' | ./logs_to_datadog_metrics") }.must_equal "foo.untagged:1|c"
  end
end
