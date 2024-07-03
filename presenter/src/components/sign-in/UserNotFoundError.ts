class UserNotFoundError extends Error {
  constructor() {
    super("User not found");
  }
}

export default UserNotFoundError