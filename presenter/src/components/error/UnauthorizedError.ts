class UnauthorizedError extends Error {
  constructor() {
    super('You are not logged in');
  }
}

export default UnauthorizedError;