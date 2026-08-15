class OrderProcessor
  def initialize(order_id)
    @order_id = order_id
  end

  def process!
    puts "Processing order ##{@order_id}"
    validate_order
    fulfill_order
  end

  private

  def validate_order
    puts "Order ##{@order_id} validated successfully."
  end

  def fulfill_order
    puts "Order ##{@order_id} fulfilled."
  end
end

processor = OrderProcessor.new("ORD-9921")
processor.process!
