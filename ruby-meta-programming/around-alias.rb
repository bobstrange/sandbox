class String
  alias_method :original_size, :size

  def size
    before_size_hook
    data = original_size
    after_size_hook
    data
  end

  private

  def before_size_hook
    p 'size start'
  end

  def after_size_hook
    p 'size finish'
  end
end

p 'Test'.size
p 'Test'.original_size
