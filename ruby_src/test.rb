# (DEAD CODE)
def abandoned_ruby_routine
  isolated_val = 999
  puts "Dead ruby code: #{isolated_val}"
end

class OrderProcessor
  def initialize(order_id)
    @order_id = order_id
    @order_count = 0
  end

  # (SOURCE)
  def fetch_untrusted_order_spec
    "ORDER_XML_<payload>PAYMENT_INJECTION</payload>"
  end

  # (SANITIZER)
  def sanitize_order_spec(raw_spec)
    raw_spec.gsub("<", "&lt;").gsub(">", "&gt;")
  end

  # (SINK)
  def dispatch_to_fulfillment_gateway(safe_spec)
    puts "[RUBY SINK GATEWAY]: Fulfilling order payload: #{safe_spec}"
  end

  def process!
    validate_order
    raw = fetch_untrusted_order_spec
    clean = sanitize_order_spec(raw)
    dispatch_to_fulfillment_gateway(clean)
  end

  private

  def validate_order
    fulfill_order
    puts "Order ##{@order_id} validated successfully."
  end

  def fulfill_order
    puts "Order ##{@order_id} fulfilled."
  end
end

processor = OrderProcessor.new("ORD-9921")
processor.process!

