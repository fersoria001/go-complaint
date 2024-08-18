const timeLeft = (date: string) => {
  const obj = new Date(parseInt(date));
  const timeLeft =
    obj.getTime() + 1000 * 60 * 60 * 24 * 5 - new Date().getTime();
  const timeLeftInDays = Math.floor(timeLeft / (1000 * 60 * 60 * 24));
  const timeLeftInHours = Math.floor(timeLeft / (1000 * 60 * 60));
  const timeLeftInMinutes = Math.floor(timeLeft / (1000 * 60));
  if (timeLeftInDays > 0) {
    return `${timeLeftInDays} days`;
  } else if (timeLeftInHours > 0) {
    return `${timeLeftInHours} hours`;
  } else {
    return `${timeLeftInMinutes} minutes`;
  }
};

export default timeLeft;
