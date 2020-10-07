// Generated from c:\gitrepos\sashimi\grammar\Sashimi.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SashimiParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		DIRECTIVE=1, COMMAND=2, ENTITY=3, SEPERATOR=4, PROP_START=5, OF=6, IS=7, 
		AS=8, LPAREN=9, RPAREN=10, HLPAREN=11, HRPAREN=12, TYPE=13, UNION=14, 
		LIST=15, CONSTRAINT=16, ALIAS=17, IDENT=18, WS=19;
	public static final int
		RULE_union_decl = 0, RULE_type_decl = 1, RULE_list_decl = 2, RULE_type_def = 3, 
		RULE_alias_decl = 4, RULE_type_is = 5, RULE_prop_decl = 6, RULE_command_call = 7, 
		RULE_entity_def = 8, RULE_export = 9;
	public static final String[] ruleNames = {
		"union_decl", "type_decl", "list_decl", "type_def", "alias_decl", "type_is", 
		"prop_decl", "command_call", "entity_def", "export"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'sashimi:'", null, "'entity'", "','", "'-'", "'of'", "'is'", "'as'", 
		"'('", "')'", "'['", "']'", null, null, "'list'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "DIRECTIVE", "COMMAND", "ENTITY", "SEPERATOR", "PROP_START", "OF", 
		"IS", "AS", "LPAREN", "RPAREN", "HLPAREN", "HRPAREN", "TYPE", "UNION", 
		"LIST", "CONSTRAINT", "ALIAS", "IDENT", "WS"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Sashimi.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SashimiParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class Union_declContext extends ParserRuleContext {
		public TerminalNode UNION() { return getToken(SashimiParser.UNION, 0); }
		public TerminalNode LPAREN() { return getToken(SashimiParser.LPAREN, 0); }
		public List<TerminalNode> ALIAS() { return getTokens(SashimiParser.ALIAS); }
		public TerminalNode ALIAS(int i) {
			return getToken(SashimiParser.ALIAS, i);
		}
		public TerminalNode RPAREN() { return getToken(SashimiParser.RPAREN, 0); }
		public List<TerminalNode> SEPERATOR() { return getTokens(SashimiParser.SEPERATOR); }
		public TerminalNode SEPERATOR(int i) {
			return getToken(SashimiParser.SEPERATOR, i);
		}
		public Union_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_union_decl; }
	}

	public final Union_declContext union_decl() throws RecognitionException {
		Union_declContext _localctx = new Union_declContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_union_decl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			match(UNION);
			setState(21);
			match(LPAREN);
			setState(22);
			match(ALIAS);
			setState(27);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPERATOR) {
				{
				{
				setState(23);
				match(SEPERATOR);
				setState(24);
				match(ALIAS);
				}
				}
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(30);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_declContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(SashimiParser.TYPE, 0); }
		public TerminalNode CONSTRAINT() { return getToken(SashimiParser.CONSTRAINT, 0); }
		public Type_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_decl; }
	}

	public final Type_declContext type_decl() throws RecognitionException {
		Type_declContext _localctx = new Type_declContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_type_decl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(32);
			match(TYPE);
			setState(34);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==CONSTRAINT) {
				{
				setState(33);
				match(CONSTRAINT);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class List_declContext extends ParserRuleContext {
		public TerminalNode LIST() { return getToken(SashimiParser.LIST, 0); }
		public Type_declContext type_decl() {
			return getRuleContext(Type_declContext.class,0);
		}
		public List_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_decl; }
	}

	public final List_declContext list_decl() throws RecognitionException {
		List_declContext _localctx = new List_declContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_list_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(36);
			match(LIST);
			setState(37);
			type_decl();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_defContext extends ParserRuleContext {
		public List_declContext list_decl() {
			return getRuleContext(List_declContext.class,0);
		}
		public Union_declContext union_decl() {
			return getRuleContext(Union_declContext.class,0);
		}
		public Type_declContext type_decl() {
			return getRuleContext(Type_declContext.class,0);
		}
		public Type_defContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_def; }
	}

	public final Type_defContext type_def() throws RecognitionException {
		Type_defContext _localctx = new Type_defContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_type_def);
		try {
			setState(42);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LIST:
				enterOuterAlt(_localctx, 1);
				{
				setState(39);
				list_decl();
				}
				break;
			case UNION:
				enterOuterAlt(_localctx, 2);
				{
				setState(40);
				union_decl();
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(41);
				type_decl();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Alias_declContext extends ParserRuleContext {
		public TerminalNode AS() { return getToken(SashimiParser.AS, 0); }
		public TerminalNode ALIAS() { return getToken(SashimiParser.ALIAS, 0); }
		public Alias_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_alias_decl; }
	}

	public final Alias_declContext alias_decl() throws RecognitionException {
		Alias_declContext _localctx = new Alias_declContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_alias_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(AS);
			setState(45);
			match(ALIAS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_isContext extends ParserRuleContext {
		public TerminalNode IS() { return getToken(SashimiParser.IS, 0); }
		public Type_defContext type_def() {
			return getRuleContext(Type_defContext.class,0);
		}
		public Type_isContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_is; }
	}

	public final Type_isContext type_is() throws RecognitionException {
		Type_isContext _localctx = new Type_isContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_type_is);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			match(IS);
			setState(48);
			type_def();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Prop_declContext extends ParserRuleContext {
		public TerminalNode PROP_START() { return getToken(SashimiParser.PROP_START, 0); }
		public TerminalNode IDENT() { return getToken(SashimiParser.IDENT, 0); }
		public Type_isContext type_is() {
			return getRuleContext(Type_isContext.class,0);
		}
		public Alias_declContext alias_decl() {
			return getRuleContext(Alias_declContext.class,0);
		}
		public Prop_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prop_decl; }
	}

	public final Prop_declContext prop_decl() throws RecognitionException {
		Prop_declContext _localctx = new Prop_declContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_prop_decl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			match(PROP_START);
			setState(51);
			match(IDENT);
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==AS) {
				{
				setState(52);
				alias_decl();
				}
			}

			setState(55);
			type_is();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Command_callContext extends ParserRuleContext {
		public TerminalNode COMMAND() { return getToken(SashimiParser.COMMAND, 0); }
		public TerminalNode LPAREN() { return getToken(SashimiParser.LPAREN, 0); }
		public TerminalNode IDENT() { return getToken(SashimiParser.IDENT, 0); }
		public TerminalNode RPAREN() { return getToken(SashimiParser.RPAREN, 0); }
		public Command_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_command_call; }
	}

	public final Command_callContext command_call() throws RecognitionException {
		Command_callContext _localctx = new Command_callContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_command_call);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(57);
			match(COMMAND);
			setState(58);
			match(LPAREN);
			setState(59);
			match(IDENT);
			setState(60);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Entity_defContext extends ParserRuleContext {
		public TerminalNode ENTITY() { return getToken(SashimiParser.ENTITY, 0); }
		public TerminalNode LPAREN() { return getToken(SashimiParser.LPAREN, 0); }
		public TerminalNode IDENT() { return getToken(SashimiParser.IDENT, 0); }
		public TerminalNode RPAREN() { return getToken(SashimiParser.RPAREN, 0); }
		public TerminalNode OF() { return getToken(SashimiParser.OF, 0); }
		public List<Prop_declContext> prop_decl() {
			return getRuleContexts(Prop_declContext.class);
		}
		public Prop_declContext prop_decl(int i) {
			return getRuleContext(Prop_declContext.class,i);
		}
		public Entity_defContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_entity_def; }
	}

	public final Entity_defContext entity_def() throws RecognitionException {
		Entity_defContext _localctx = new Entity_defContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_entity_def);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(ENTITY);
			setState(63);
			match(LPAREN);
			setState(64);
			match(IDENT);
			setState(65);
			match(RPAREN);
			setState(66);
			match(OF);
			setState(67);
			prop_decl();
			setState(71);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PROP_START) {
				{
				{
				setState(68);
				prop_decl();
				}
				}
				setState(73);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExportContext extends ParserRuleContext {
		public TerminalNode DIRECTIVE() { return getToken(SashimiParser.DIRECTIVE, 0); }
		public Command_callContext command_call() {
			return getRuleContext(Command_callContext.class,0);
		}
		public Entity_defContext entity_def() {
			return getRuleContext(Entity_defContext.class,0);
		}
		public ExportContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_export; }
	}

	public final ExportContext export() throws RecognitionException {
		ExportContext _localctx = new ExportContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_export);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			match(DIRECTIVE);
			setState(77);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case COMMAND:
				{
				setState(75);
				command_call();
				}
				break;
			case ENTITY:
				{
				setState(76);
				entity_def();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\25R\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t\13\3"+
		"\2\3\2\3\2\3\2\3\2\7\2\34\n\2\f\2\16\2\37\13\2\3\2\3\2\3\3\3\3\5\3%\n"+
		"\3\3\4\3\4\3\4\3\5\3\5\3\5\5\5-\n\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3"+
		"\b\5\b8\n\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\7"+
		"\nH\n\n\f\n\16\nK\13\n\3\13\3\13\3\13\5\13P\n\13\3\13\2\2\f\2\4\6\b\n"+
		"\f\16\20\22\24\2\2\2N\2\26\3\2\2\2\4\"\3\2\2\2\6&\3\2\2\2\b,\3\2\2\2\n"+
		".\3\2\2\2\f\61\3\2\2\2\16\64\3\2\2\2\20;\3\2\2\2\22@\3\2\2\2\24L\3\2\2"+
		"\2\26\27\7\20\2\2\27\30\7\13\2\2\30\35\7\23\2\2\31\32\7\6\2\2\32\34\7"+
		"\23\2\2\33\31\3\2\2\2\34\37\3\2\2\2\35\33\3\2\2\2\35\36\3\2\2\2\36 \3"+
		"\2\2\2\37\35\3\2\2\2 !\7\f\2\2!\3\3\2\2\2\"$\7\17\2\2#%\7\22\2\2$#\3\2"+
		"\2\2$%\3\2\2\2%\5\3\2\2\2&\'\7\21\2\2\'(\5\4\3\2(\7\3\2\2\2)-\5\6\4\2"+
		"*-\5\2\2\2+-\5\4\3\2,)\3\2\2\2,*\3\2\2\2,+\3\2\2\2-\t\3\2\2\2./\7\n\2"+
		"\2/\60\7\23\2\2\60\13\3\2\2\2\61\62\7\t\2\2\62\63\5\b\5\2\63\r\3\2\2\2"+
		"\64\65\7\7\2\2\65\67\7\24\2\2\668\5\n\6\2\67\66\3\2\2\2\678\3\2\2\289"+
		"\3\2\2\29:\5\f\7\2:\17\3\2\2\2;<\7\4\2\2<=\7\13\2\2=>\7\24\2\2>?\7\f\2"+
		"\2?\21\3\2\2\2@A\7\5\2\2AB\7\13\2\2BC\7\24\2\2CD\7\f\2\2DE\7\b\2\2EI\5"+
		"\16\b\2FH\5\16\b\2GF\3\2\2\2HK\3\2\2\2IG\3\2\2\2IJ\3\2\2\2J\23\3\2\2\2"+
		"KI\3\2\2\2LO\7\3\2\2MP\5\20\t\2NP\5\22\n\2OM\3\2\2\2ON\3\2\2\2P\25\3\2"+
		"\2\2\b\35$,\67IO";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}