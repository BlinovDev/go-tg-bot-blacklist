# go-tg-bot-blacklist

## Description

This package is a lightweight solution to manage the blacklist for your telegram bot.

## Integration

Include the package:
```Go
tgblacklist "github.com/BlinovDev/go-tg-bot-blacklist"
```

In your code set strategy and use public methods:
```Go
// SET STRATEGY
tgblacklist.SetStrategy("BL")

// ADD USER TO LIST
tgblacklist.AddToList("BotFather") // no @ in the beginning

// CHECK IF USER CAN CONTINUE
username := message.From.UserName
blocked, _ := tgblacklist.IsBlocked(username)
if blocked {
	fmt.Printf("User %s is blocked! Continue...", username)
	continue
}
```

## Main features(will be marked as soon as done)

- [x] Check if user in a BL by user_name;
- [x] Add user to BL;
- [ ] Return user_names from BL;
- [ ] Delete user from BL after set time(optional);
- [ ] Describe all public methods for better user experience;
