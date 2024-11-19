# curl -X PUT -d "title=Updated%20Egyptian%20Hieroglyph&slug=updated-egyptian-heiroglyph&value=9500.75&description=An%20updated%20ancient%20Egyptian%20writing&details=Found%20in%20the%20Egyptian%20desert%20and%20estimated%20to%20be%202100%20years%20old" localhost:9000/pieces/1

curl -X PUT -H "Content-Type: application/json" -d '{
  "title": "Updated Egyptian Hieroglyph",
  "slug": "updated-egyptian-heiroglyph",
  "value": 9500.75,
  "description": "An updated ancient Egyptian writing",
  "details": "Found in the Egyptian desert and estimated to be 2100 years old"
}' localhost:9000/pieces/3
