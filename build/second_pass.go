package build

type secondPass struct {
	*BaseSashimiListener
}

func (s *secondPass) EnterLoopCall(ctx *LoopCallContext) {}

func (s *secondPass) ExitLoopCall(ctx *LoopCallContext) {}

func (s *secondPass) EnterCommandCall(ctx *CommandCallContext) {}

func (s *secondPass) ExitCommandCall(ctx *CommandCallContext) {}

func (s *secondPass) EnterBlockScope(ctx *BlockScopeContext) {}

func (s *secondPass) ExitBlockScope(ctx *BlockScopeContext) {}
