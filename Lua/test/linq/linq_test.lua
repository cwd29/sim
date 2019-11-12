local lu = require('luaunit')

package.path = package.path..';..\\?.lua'
local linq = require('core.linq.linq')

function test_new_Table()
  linq({1}):print()

  linq({2}):print()
end

os.exit(lu.LuaUnit.run())
