extends Node

@onready var go_node = $MyGoNode

func _ready():
	# Debug signal list
	for signal_info in go_node.get_signal_list():
		print("Signal:", signal_info.name)

	# Now connect
	go_node.connect("my_signal_event_handler", Callable(go_node, "my_signal_handler"))
	go_node.connect("my_signal_with_params_handler", Callable(go_node, "my_signal_with_params_handler"))

	# Emit them
	go_node.emit_signal("my_signal_event_handler")
	go_node.emit_signal("my_signal_with_params_handler", "Hello from GDScript!", 10)
