# Various components and hooks in the GUI are not pure and may not be testable.
# The logout() function doesn't re render at first, you have to click 2 times to see the result.
# In the complaint page when the chat is closed it's not closing the websocket connection. (done)
# When you retrieve hireable users from an enterprise, it still returns users that have been already
hired in the current enterprise.(done)
# The send complaint GUI is not correctly displaying error messages.
# Whenever a user is hired, it start as an assistant, it does not reflect the GUI, it needs to be approved
with the correct employ position. (done)
# When you update your personal settings, all your complaints and messages should reflect that update,
with the current domain logic, you may need to update one by one each piece of data.(done)
# When is the owner who ask for a review of the complaint, the assistant doesnt exist because its ID is created
in the MoveToReview from the go graphql schema.(done)
# The chat-handler, room and client started to grow, it has a few race conditions that doesn't seem to be a 
problem right now but it impedes the constant grow of the feature, it may need to be refactored and appended to the
domain layer.(done)
