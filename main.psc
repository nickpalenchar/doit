import ./hooks/*

class Parser {

  constructor() {
    self.hooks = []
  }
  fn register_hook(hook: Hook) {
    self.hooks.append(hook)
  }
  fn exec(json: obj) {
    for hook in self.hooks {
      if hook.required && not (hook.name in keys(json) && union(hook.will_accept, keys(json))) {
        throw Error("Parse Error: Missing required hook name ${hook.name}")
    order = sort(self.hooks, fn(prev, hook) { return hook.priority < prev.priority }, {priority: -999... })
    
    section = json[hook.name]
     
    for hook in order {
      if premain in hook {
        
    


class Hook {
  name: string // name of the hook, will match the yml top level, eg '__MAIN__'
  will_accept: string[] // aliases of hooks that can also be accepted. Can be dynamic. A regexp can be used and the subgroups are passed as args.
	           // doit is meant to be static so this should be used sparingly. In general, it is reserved for future use and there is not yet a good use case for it
  
  priority: int // order to run. Recommended Multiples of 10 for easy insertion, then 5, then 1.
  required: bool // if true, will error if not present. This should only be used for '__MAIN__'
  
  fn is_optional() { return not self.required }

  abstract fn premain(section, json) // execute before main. first arg is the entire value within the hook, second is the entire json of the file.
  abstract fn main(section)
  abstract fn postmain(section, json) // after main.
}

function main {

  
  for hook in hooks {
    parser.register_hook(hook)
  }

  text = readfile(system.argv[2])
  formatted = json.parse(text)

  parser.exec(formatted)

}
