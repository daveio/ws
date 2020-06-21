package data

const HeaderComment string = "# BEGIN ws env integration"

const FooterComment string = "# END ws env integration"

const BashTemplate string = `
WS_OLDPROMPT="${PROMPT_COMMAND}"
PROMPT_COMMAND='eval $(ws env)'
PROMPT_COMMAND+="; ${WS_OLDPROMPT}"
unset WS_OLDPROMPT
`

const ZshTemplate string = `
ws_hook () {
	eval "$(ws env)"
}
precmd_functions+=ws_hook
`

const FishTemplate string = `
function ws_hook --on-event fish_prompt
  set ws_vars (ws env)
  for var in $ws_vars
    set pair (string split = $var)
    set -gx $pair[1] $pair[2]
  end
end
`
