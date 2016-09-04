user = box.space.users
if not user then
    user = box.schema.create_space('users')
    user:create_index('primary')
end
function tnt ()
    return "Hello! :)"
end