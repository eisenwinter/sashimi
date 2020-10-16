package build

type secondPass struct {
	*BaseSashimiListener
}

/*
	Hookay basically here is what we need to do:
	First - process the layout pages (layout_section), generate those fully and ready in some intermediate format
	Second - walk the files, if we encounter a sashimi:layout directive we have the layout_section ready to inline the layout
	For each file
		- check if there are any command calls on entities that are not within a repeat scope, those pages will be duplicated for each entity
		- once we know wheter a page is a single page or is going to be multiple pages we start walking the dom tree
		- for each loop directive we extract the whole sccope as fragment
		- for each link directive we look ahead for the next a tag (the next tag SHOULD be an a tag)
		- for each display directive we try to resolve the entrity and replace any block content that currently exists with the resolved content

	for nested repeats we resolve inner to most outter scope
	we can basically discard any entity definitions at this point - all we need are command and loop calls

	well this should be easy peasy lemon squeezy - but i guess its going to be hard hard lemon hard.

	further we could store the offset within the html of each call as well as scope levels in the analyze run to get a better understanding
	but idk if thats necassary we could as well just get offsets here tbd
*/

func (s *secondPass) EnterLoopCall(ctx *LoopCallContext) {}

func (s *secondPass) ExitLoopCall(ctx *LoopCallContext) {}

func (s *secondPass) EnterCommandCall(ctx *CommandCallContext) {}

func (s *secondPass) ExitCommandCall(ctx *CommandCallContext) {}

func (s *secondPass) EnterBlockScope(ctx *BlockScopeContext) {}

func (s *secondPass) ExitBlockScope(ctx *BlockScopeContext) {}
