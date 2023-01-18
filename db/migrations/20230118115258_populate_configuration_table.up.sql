INSERT INTO configuations (id, configuation, "role", created_at, updated_at)
  VALUES('1', '{
  "items": {
    "resources": {
      "list": {
        "action": "allow"
      }
    },
    "default": "disallow"
  },
  "cart": {
    "default": "allow"
  },
  "accounts": {
    "default": "disallow"
  }
}', 'user', now(), now()
);


INSERT INTO configuations (id, configuation, "role", created_at, updated_at)
  VALUES('2', '{
    "items": {
      "resources": {
        "add": {
          "action": "allow"
        }
      },
      "default": "disallow"
    },
    "accounts": {
      "resources": {
        "suspend": {
          "action": "allow"
        }
      },
      "default": "disallow"
    }
  }', 'admin', now(), now()
);

