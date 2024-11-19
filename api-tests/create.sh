curl -X POST -H "Content-Type: application/json" -d '{
  "title": "Egyptian heiroglyph",
  "slug": "egyptian-heiroglyph",
  "value": 9000.50,
  "description": "An ancient Egyptian writing",
  "details": "Discovered in the Egyptian desert off the coast of the Mediterranean. Estimated to be approximately 2000 years old"
}' localhost:9000/pieces

curl -X POST -H "Content-Type: application/json" -d '{
  "title": "Kenyan Oduor",
  "slug": "kenyan-oduor",
  "value": 20300123030.50,
  "description": "An ancient Kenyan human skull",
  "details": "Discovered in the Kenyan arid area of Turkana next to Lake Turakana. Estimated to be approximately 2004000 years old"
}' localhost:9000/pieces
