package main

// ProcessPluginRPCServer ProcessPluginRPCServer 정의
type ProcessPluginRPCServer struct{ Impl interface{} }

// Hello RPC 테스트 함수
func (s *ProcessPluginRPCServer) Hello(args interface{}, resp *string) error {
	impl := s.Impl.(*Plugin)
	*resp = impl.Hello()
	return nil
}
