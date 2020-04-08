extends KinematicBody2D

const UP = Vector2(0, -1)
const MAX_SPEED = 300
const GRAVITY = 20
const ACCELARATION = 50
const JUMP_HEIGHT = -800

var motion = Vector2()

func _physics_process(delta):
	motion.y += GRAVITY
	var friction = false

	if Input.is_action_pressed("ui_right"):
		motion.x = min(motion.x + ACCELARATION, MAX_SPEED)
		$Sprite.flip_h = false
		$Sprite.play("Run")
		
	elif Input.is_action_pressed("ui_left"):
		motion.x = max(motion.x - ACCELARATION, -MAX_SPEED)
		$Sprite.flip_h = true
		$Sprite.play("Run")
		
	else:
		friction = true
		$Sprite.play("Idle")
		
	if is_on_floor():
		if Input.is_action_just_pressed("ui_up"):
			motion.y = JUMP_HEIGHT
		if friction == true:
			motion.x = lerp(motion.x, 0, 0.15)
			
	else:
		if motion.y < 0:
			$Sprite.play("Jump")
		else:
			if $Sprite.playing:
				$Sprite.play("Fall")
			else:
				$Sprite.play("Run")			
		if Input.is_action_just_pressed("ui_down"):
			motion.y = -JUMP_HEIGHT/4
		if friction == true:
			motion.x = lerp(motion.x, 0, 0.1)
			
		
	motion = move_and_slide(motion, UP)
