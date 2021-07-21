const unauthenticatedResponse = (res, message) => {
  res.status(401).json({
    message
  });
}

module.exports = {
  unauthenticatedResponse
}