package webserver

// Telemetry represents a single UDP message from the telemetry system.
// For fields like Acceleration, X = right, Y = up, Z = forward.
// For fields like AngularVelocity, X = pitch, Y = yaw, Z = roll.
type Telemetry struct {
	IsRaceOn         int32   `json:"is_race_on"`
	TimestampMS      uint32  `json:"timestamp_ms"`
	EngineMaxRpm     float32 `json:"engine_max_rpm"`
	EngineIdleRpm    float32 `json:"engine_idle_rpm"`
	CurrentEngineRpm float32 `json:"current_engine_rpm"`

	AccelerationX float32 `json:"acceleration_x"` // In the car's local space; X = right, Y = up, Z = forward
	AccelerationY float32 `json:"acceleration_y"`
	AccelerationZ float32 `json:"acceleration_z"`

	VelocityX float32 `json:"velocity_x"` // In the car's local space; X = right, Y = up, Z = forward
	VelocityY float32 `json:"velocity_y"`
	VelocityZ float32 `json:"velocity_z"`

	AngularVelocityX float32 `json:"angular_velocity_x"` // In the car's local space; X = pitch, Y = yaw, Z = roll
	AngularVelocityY float32 `json:"angular_velocity_y"`
	AngularVelocityZ float32 `json:"angular_velocity_z"`

	Yaw   float32 `json:"yaw"`
	Pitch float32 `json:"pitch"`
	Roll  float32 `json:"roll"`

	NormalizedSuspensionTravelFrontLeft  float32 `json:"normalized_suspension_travel_front_left"` // Suspension travel normalized: 0.0f = max stretch; 1.0 = max compression
	NormalizedSuspensionTravelFrontRight float32 `json:"normalized_suspension_travel_front_right"`
	NormalizedSuspensionTravelRearLeft   float32 `json:"normalized_suspension_travel_rear_left"`
	NormalizedSuspensionTravelRearRight  float32 `json:"normalized_suspension_travel_rear_right"`

	TireSlipRatioFrontLeft  float32 `json:"tire_slip_ratio_front_left"` // Tire normalized slip ratio, 0 means 100% grip and |ratio| > 1.0 means loss of grip
	TireSlipRatioFrontRight float32 `json:"tire_slip_ratio_front_right"`
	TireSlipRatioRearLeft   float32 `json:"tire_slip_ratio_rear_left"`
	TireSlipRatioRearRight  float32 `json:"tire_slip_ratio_rear_right"`

	WheelRotationSpeedFrontLeft  float32 `json:"wheel_rotation_speed_front_left"` // Wheel rotation speed in radians/sec.
	WheelRotationSpeedFrontRight float32 `json:"wheel_rotation_speed_front_right"`
	WheelRotationSpeedRearLeft   float32 `json:"wheel_rotation_speed_rear_left"`
	WheelRotationSpeedRearRight  float32 `json:"wheel_rotation_speed_rear_right"`

	WheelOnRumbleStripFrontLeft  int32 `json:"wheel_on_rumble_strip_front_left"` // = 1 when wheel is on rumble strip, = 0 when off.
	WheelOnRumbleStripFrontRight int32 `json:"wheel_on_rumble_strip_front_right"`
	WheelOnRumbleStripRearLeft   int32 `json:"wheel_on_rumble_strip_rear_left"`
	WheelOnRumbleStripRearRight  int32 `json:"wheel_on_rumble_strip_rear_right"`

	WheelInPuddleDepthFrontLeft  float32 `json:"wheel_in_puddle_depth_front_left"` // = from 0 to 1, where 1 is the deepest puddle.
	WheelInPuddleDepthFrontRight float32 `json:"wheel_in_puddle_depth_front_right"`
	WheelInPuddleDepthRearLeft   float32 `json:"wheel_in_puddle_depth_rear_left"`
	WheelInPuddleDepthRearRight  float32 `json:"wheel_in_puddle_depth_rear_right"`

	SurfaceRumbleFrontLeft  float32 `json:"surface_rumble_front_left"` // Non-dimensional surface rumble values passed to controller force feedback
	SurfaceRumbleFrontRight float32 `json:"surface_rumble_front_right"`
	SurfaceRumbleRearLeft   float32 `json:"surface_rumble_rear_left"`
	SurfaceRumbleRearRight  float32 `json:"surface_rumble_rear_right"`

	TireSlipAngleFrontLeft  float32 `json:"tire_slip_angle_front_left"` // Tire normalized slip angle, = 0 means 100% grip and |angle| > 1.0 means loss of grip
	TireSlipAngleFrontRight float32 `json:"tire_slip_angle_front_right"`
	TireSlipAngleRearLeft   float32 `json:"tire_slip_angle_rear_left"`
	TireSlipAngleRearRight  float32 `json:"tire_slip_angle_rear_right"`

	TireCombinedSlipFrontLeft  float32 `json:"tire_combined_slip_front_left"` // Tire normalized combined slip, = 0 means 100% grip and |slip| > 1.0 means loss of grip
	TireCombinedSlipFrontRight float32 `json:"tire_combined_slip_front_right"`
	TireCombinedSlipRearLeft   float32 `json:"tire_combined_slip_rear_left"`
	TireCombinedSlipRearRight  float32 `json:"tire_combined_slip_rear_right"`

	SuspensionTravelMetersFrontLeft  float32 `json:"suspension_travel_meters_front_left"` // Actual suspension travel in meters
	SuspensionTravelMetersFrontRight float32 `json:"suspension_travel_meters_front_right"`
	SuspensionTravelMetersRearLeft   float32 `json:"suspension_travel_meters_rear_left"`
	SuspensionTravelMetersRearRight  float32 `json:"suspension_travel_meters_rear_right"`

	CarOrdinal          int32 `json:"car_ordinal"`           // Unique ID of the car make/model
	CarClass            int32 `json:"car_class"`             // Between 0 (D -- worst cars) and 7 (X class -- best cars) inclusive
	CarPerformanceIndex int32 `json:"car_performance_index"` // Between 100 (slowest car) and 999 (fastest car) inclusive
	DrivetrainType      int32 `json:"drivetrain_type"`       // Corresponds to EDrivetrainType; 0 = FWD, 1 = RWD, 2 = AWD
	NumCylinders        int32 `json:"num_cylinders"`         // Number of cylinders in the engine

	// The following are part of V2.   They are likely missing from Forza 7 if the data out type is set to "SLED"
	PositionX float32 `json:"position_x"` // Position (meters)
	PositionY float32 `json:"position_y"`
	PositionZ float32 `json:"position_z"`

	Speed  float32 `json:"speed"`  // meters / sec
	Power  float32 `json:"power"`  // watts
	Torque float32 `json:"torque"` // newton meters

	TireTempFrontLeft  float32 `json:"tire_time_front_left"`
	TireTempFrontRight float32 `json:"tire_time_front_right"`
	TireTempRearLeft   float32 `json:"tire_time_rear_left"`
	TireTempRearRight  float32 `json:"tire_time_rear_right"`

	Boost            float32 `json:"boost"`
	Fuel             float32 `json:"fuel"`
	DistanceTraveled float32 `json:"distance_traveled"`
	BestLap          float32 `json:"best_lap"`
	LastLap          float32 `json:"last_lap"`
	CurrentLap       float32 `json:"current_lap"`
	CurrentRaceTime  float32 `json:"current_race_time"`

	LapNumber    uint16 `json:"lap_number"`
	RacePosition uint8  `json:"race_position"`

	Accel     uint8 `json:"accel"`
	Brake     uint8 `json:"brake"`
	Clutch    uint8 `json:"clutch"`
	HandBrake uint8 `json:"hand_brake"`
	Gear      uint8 `json:"gear"`
	Steer     int8  `json:"steer"`

	NormalizedDrivingLine       int8 `json:"normalized_driving_line"`
	NormalizedAIBrakeDifference int8 `json:"normalized_ai_brake_difference"`
}
