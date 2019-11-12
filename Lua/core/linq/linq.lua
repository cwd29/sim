local mt = {}
local index = {}


function index:each(action)
  for k, v in pairs(self.db) do

  end
end

function index:value()
  print('this is a test');
end

function index:print()
  local result = '['
  for i, k in pairs(self.db) do
    print(i ..' --'..k)
  end
end

function index:where(action)
  table.insert(self.chain, function(k, v)
    if action(k, v) then
      index:next(k, v)
    end
  end)
end

mt.__index = index
return setmetatable({}, {
  __call = function(m, newtab)
    if type(newtab) ~= 'table' then
      error('the parameter what you passed in is not a table', 2)
    end

    return setmetatable({db = newtab , chain = {}}, mt)
  end
})
