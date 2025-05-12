schema "public" {}

enum "event_attendance_status_enum" {
	schema = schema.public
	values = [
		"GOING",
		"MAYBE",
		"NOT_GOING",
	]
}
enum "event_confidentiality_enum" {
	schema = schema.public
	values = [
		"PUBLIC",
		"ONLY_INVITED",
	]
}
enum "kind_enum" {
	schema = schema.public
	values = [
		"RPG",
		"GAME",
		"EXPANSION",
	]
}
enum "otp_purpose_enum" {
	schema = schema.public
	values = [
		"SIGN_IN",
	]
}
enum "provider_enum" {
	schema = schema.public
	values = [
		"GOOGLE",
		"LUDOPEDIA",
	]
}


table "accounts" {
	schema = schema.public
	column "avatar_path" {
		type = sql("VARCHAR(128)")
		null = true
	}
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "handle" {
		type = sql("VARCHAR(16)")
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	column "is_admin" {
		type = sql("BOOLEAN")
		default = sql("FALSE")
	}
	column "name" {
		type = sql("VARCHAR(64)")
		null = true
	}
	primary_key {
		columns = [
			column.id,
		]
	}
	index "accounts_handle_idx" {
		columns = [
			column.handle,
		]
		unique = true
	}
}
table "connections" {
	schema = schema.public
	column "access_token" {
		type = sql("VARCHAR(500)")
		null = true
	}
	column "account_id" {
		type = sql("INTEGER")
	}
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "external_handle" {
		type = sql("VARCHAR(30)")
		null = true
	}
	column "external_id" {
		type = sql("VARCHAR(255)")
	}
	column "provider" {
		type = enum.provider_enum
	}
	column "refresh_token" {
		type = sql("VARCHAR(500)")
		null = true
	}
	primary_key {
		columns = [
			column.external_id,
			column.provider,
		]
	}
	index "connections_account_id_idx" {
		columns = [
			column.account_id,
		]
	}
	foreign_key "connections_account_id_fk" {
		columns = [
			column.account_id
		]
		ref_columns = [
			table.accounts.column.id
		]
		on_delete = CASCADE
	}
}
table "email_addresses" {
	schema = schema.public
	column "account_id" {
		type = sql("INTEGER")
	}
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "email_address" {
		type = sql("VARCHAR(500)")
	}
	column "validated_at" {
		type = sql("TIMESTAMPTZ")
		null = true
	}
	primary_key {
		columns = [
			column.email_address,
		]
	}
	foreign_key "email_addresses_account_id_fk" {
		columns = [
			column.account_id
		]
		ref_columns = [
			table.accounts.column.id
		]
		on_delete = CASCADE
	}
}
table "event_attendances" {
	schema = schema.public
	column "account_id" {
		type = sql("INTEGER")
	}
	column "confirmed_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "event_id" {
		type = sql("INTEGER")
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	column "status" {
		type = enum.event_attendance_status_enum
	}
	primary_key {
		columns = [
			column.account_id,
			column.event_id,
			column.id,
		]
	}
	foreign_key "event_attendances_event_id_fk" {
		columns = [
			column.event_id
		]
		ref_columns = [
			table.events.column.id
		]
	}
	foreign_key "event_attendances_account_id_fk" {
		columns = [
			column.account_id
		]
		ref_columns = [
			table.accounts.column.id
		]
	}
}
table "event_games" {
	schema = schema.public
	column "event_id" {
		type = sql("INTEGER")
	}
	column "game_id" {
		type = sql("INTEGER")
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	primary_key {
		columns = [
			column.event_id,
			column.game_id,
			column.id,
		]
	}
	foreign_key "event_games_event_id_fk" {
		columns = [
			column.event_id
		]
		ref_columns = [
			table.events.column.id
		]
	}
	foreign_key "event_games_game_id_fk" {
		columns = [
			column.game_id
		]
		ref_columns = [
			table.games.column.id
		]
	}
}
table "events" {
	schema = schema.public
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "date" {
		type = sql("TIMESTAMPTZ")
	}
	column "description" {
		type = sql("VARCHAR(1000)")
	}
	column "icon_path" {
		type = sql("VARCHAR(250)")
		null = true
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	column "location_address" {
		type = sql("VARCHAR(500)")
	}
	column "location_name" {
		type = sql("VARCHAR(100)")
	}
	column "max_amount_of_players" {
		type = sql("INTEGER")
		null = true
	}
	column "name" {
		type = sql("VARCHAR(50)")
	}
	column "owner_id" {
		type = sql("INTEGER")
	}
	primary_key {
		columns = [
			column.id,
		]
	}
	foreign_key "events_owner_id_fk" {
		columns = [
			column.owner_id
		]
		ref_columns = [
			table.accounts.column.id
		]
	}
}
table "games" {
	schema = schema.public
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "description" {
		type = sql("VARCHAR(1000)")
	}
	column "icon_path" {
		type = sql("VARCHAR(256)")
		null = true
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	column "kind" {
		type = enum.kind_enum
	}
	column "ludopedia_id" {
		type = sql("INTEGER")
		null = true
	}
	column "ludopedia_url" {
		type = sql("VARCHAR(512)")
		null = true
	}
	column "max_amount_of_players" {
		type = sql("INTEGER")
	}
	column "min_amount_of_players" {
		type = sql("INTEGER")
	}
	column "name" {
		type = sql("VARCHAR(128)")
	}
	primary_key {
		columns = [
			column.id,
		]
	}
	index "games_ludopedia_id_idx" {
		columns = [
			column.ludopedia_id,
		]
	}
}
table "medias" {
	schema = schema.public
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "date" {
		type = sql("TIMESTAMPTZ")
	}
	column "description" {
		type = sql("VARCHAR(500)")
		null = true
	}
	column "game_id" {
		type = sql("INTEGER")
		null = true
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	column "owner_id" {
		type = sql("INTEGER")
	}
	column "path" {
		type = sql("VARCHAR(500)")
	}
	primary_key {
		columns = [
			column.id,
		]
	}
	foreign_key "medias_owner_id_fk" {
		columns = [
			column.owner_id
		]
		ref_columns = [
			table.accounts.column.id
		]
	}
}
table "one_time_passwords" {
	schema = schema.public
	column "account_id" {
		type = sql("INTEGER")
	}
	column "code" {
		type = sql("VARCHAR(255)")
	}
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "purpose" {
		type = enum.otp_purpose_enum
	}
	primary_key {
		columns = [
			column.account_id,
			column.code,
			column.purpose,
		]
	}
	foreign_key "one_time_passwords_account_id_fk" {
		columns = [
			column.account_id
		]
		ref_columns = [
			table.accounts.column.id
		]
		on_delete = CASCADE
	}
}
table "personal_collections" {
	schema = schema.public
	column "account_id" {
		type = sql("INTEGER")
	}
	column "acquired_at" {
		type = sql("TIMESTAMPTZ")
		null = true
	}
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "game_id" {
		type = sql("INTEGER")
	}
	column "id" {
		type = sql("INTEGER")
    identity {
			generated = ALWAYS
			start = 0
			increment = 1
    }
	}
	column "paid" {
		type = sql("INTEGER")
		null = true
	}
	primary_key {
		columns = [
			column.id,
		]
	}
	foreign_key "personal_collections_account_id_fk" {
		columns = [
			column.account_id
		]
		ref_columns = [
			table.accounts.column.id
		]
	}
	foreign_key "personal_collections_game_id_fk" {
		columns = [
			column.game_id
		]
		ref_columns = [
			table.games.column.id
		]
	}
}
table "sessions" {
	schema = schema.public
	column "account_id" {
		type = sql("INTEGER")
	}
	column "created_at" {
		type = sql("TIMESTAMPTZ")
		default = sql("NOW()")
	}
	column "session_id" {
		type = sql("CHAR(128)")
	}
	primary_key {
		columns = [
			column.session_id,
		]
	}
	foreign_key "sessions_account_id_fk" {
		columns = [
			column.account_id
		]
		ref_columns = [
			table.accounts.column.id
		]
		on_delete = CASCADE
	}
}
