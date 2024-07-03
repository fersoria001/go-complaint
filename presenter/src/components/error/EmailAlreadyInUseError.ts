class EmailAlreadyInUseError extends Error {
    constructor() {
        super("Email already in use");
    }
}

export default EmailAlreadyInUseError;