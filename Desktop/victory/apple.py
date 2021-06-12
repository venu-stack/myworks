from hashlib import sha256

first_hash = 'this is a new files filling'

hash_out = sha256(first_hash.encode()).hexdigest()

print(hash_out)