package data

const HeaderComment string = "# BEGIN ws env integration"

const FooterComment string = "# END ws env integration"

const BashTemplate string = `
PROMPT_COMMAND="go run . env; ${PROMPT_COMMAND}"
`

const ZshTemplate string = `
ws_hook () {
	eval $(go run . env)
}
precmd_functions+=ws_hook
`

const FishTemplate string = `
function ws_hook --on-event fish_prompt
    eval (go run . env)
end
`
