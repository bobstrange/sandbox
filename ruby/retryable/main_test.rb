require "minitest/autorun"
require "logger"

require "bundler/setup"
require "retryable"


class TestRetryable < Minitest::Test
  def log_method
    @log_method ||= lambda do |retries, exception|
      Logger.new(STDOUT).info("Attempt #{retries}, error: #{exception}")
    end
  end

  def raise_error
    raise "100% error"
  end

  def possibly_error
    if rand > 0.9
      raise "possibly error"
    end
  end

  def test_error
    error = assert_raises RuntimeError do
      Retryable.retryable(
        on:         StandardError,
        sleep:      0.1,
        tries:      3,
        log_method: log_method

      ) do
        raise_error
      end

      assert_equal("100% error", error.message)
    end
  end

  def test_error_with_ensure
    ensure_callback = lambda do |retries|
      Logger.new(STDOUT).info("total attempts: #{retries}")
    end

    error = assert_raises RuntimeError do
      Retryable.retryable(
        on:         StandardError,
        sleep:      0.1,
        tries:      3,
        log_method: log_method,
        ensure:  ensure_callback
      ) do
        raise_error
      end
    end

    # Even though we put ensure block, the error will be thrown if retries exhausted
    assert_equal("100% error", error.message)
  end
end
