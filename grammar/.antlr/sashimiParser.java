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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, DIRECTIVE=11, COMMAND=12, IDENT=13, ALIAS=14, CONSTRAINT=15, 
		SEPERATOR=16, TYPE=17, UNION=18, WS=19;
	public static final int
		RULE_union_decl = 0, RULE_type_decl = 1, RULE_list_decl = 2, RULE_type_def = 3, 
		RULE_alias_decl = 4, RULE_type_is = 5, RULE_prop_decl = 6, RULE_entity_decl = 7, 
		RULE_export = 8;
	public static final String[] ruleNames = {
		"union_decl", "type_decl", "list_decl", "type_def", "alias_decl", "type_is", 
		"prop_decl", "entity_decl", "export"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'('", "'\"'", "')'", "'['", "']'", "'list of'", "'as'", "'is'", 
		"'-'", "'of'", "'sashimi:'", null, null, null, null, "','"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, null, "DIRECTIVE", 
		"COMMAND", "IDENT", "ALIAS", "CONSTRAINT", "SEPERATOR", "TYPE", "UNION", 
		"WS"
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
		public List<TerminalNode> ALIAS() { return getTokens(SashimiParser.ALIAS); }
		public TerminalNode ALIAS(int i) {
			return getToken(SashimiParser.ALIAS, i);
		}
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
			setState(18);
			match(UNION);
			setState(19);
			match(T__0);
			setState(20);
			match(T__1);
			setState(21);
			match(ALIAS);
			setState(22);
			match(T__1);
			setState(29);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPERATOR) {
				{
				{
				setState(23);
				match(SEPERATOR);
				setState(24);
				match(T__1);
				setState(25);
				match(ALIAS);
				setState(26);
				match(T__1);
				}
				}
				setState(31);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(32);
			match(T__2);
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
			setState(34);
			match(TYPE);
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(35);
				match(T__3);
				setState(36);
				match(CONSTRAINT);
				setState(37);
				match(T__4);
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
			setState(40);
			match(T__5);
			setState(41);
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
			setState(46);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__5:
				enterOuterAlt(_localctx, 1);
				{
				setState(43);
				list_decl();
				}
				break;
			case UNION:
				enterOuterAlt(_localctx, 2);
				{
				setState(44);
				union_decl();
				}
				break;
			case TYPE:
				enterOuterAlt(_localctx, 3);
				{
				setState(45);
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
			setState(48);
			match(T__6);
			setState(49);
			match(T__1);
			setState(50);
			match(ALIAS);
			setState(51);
			match(T__1);
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
			setState(53);
			match(T__7);
			setState(54);
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
		public TerminalNode IDENT() { return getToken(SashimiParser.IDENT, 0); }
		public Type_defContext type_def() {
			return getRuleContext(Type_defContext.class,0);
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
			setState(56);
			match(T__8);
			setState(57);
			match(IDENT);
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(58);
				alias_decl();
				}
			}

			setState(61);
			match(T__7);
			setState(62);
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

	public static class Entity_declContext extends ParserRuleContext {
		public List<Prop_declContext> prop_decl() {
			return getRuleContexts(Prop_declContext.class);
		}
		public Prop_declContext prop_decl(int i) {
			return getRuleContext(Prop_declContext.class,i);
		}
		public Entity_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_entity_decl; }
	}

	public final Entity_declContext entity_decl() throws RecognitionException {
		Entity_declContext _localctx = new Entity_declContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_entity_decl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			match(T__9);
			setState(66); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(65);
				prop_decl();
				}
				}
				setState(68); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__8 );
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
		public TerminalNode COMMAND() { return getToken(SashimiParser.COMMAND, 0); }
		public TerminalNode IDENT() { return getToken(SashimiParser.IDENT, 0); }
		public List<Prop_declContext> prop_decl() {
			return getRuleContexts(Prop_declContext.class);
		}
		public Prop_declContext prop_decl(int i) {
			return getRuleContext(Prop_declContext.class,i);
		}
		public ExportContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_export; }
	}

	public final ExportContext export() throws RecognitionException {
		ExportContext _localctx = new ExportContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_export);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			match(DIRECTIVE);
			setState(71);
			match(COMMAND);
			setState(72);
			match(T__0);
			setState(73);
			match(IDENT);
			setState(74);
			match(T__2);
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__8) {
				{
				{
				setState(75);
				prop_decl();
				}
				}
				setState(80);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\25T\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\7\2\36\n\2\f\2\16\2!\13\2\3\2\3\2\3\3\3\3\3\3"+
		"\3\3\5\3)\n\3\3\4\3\4\3\4\3\5\3\5\3\5\5\5\61\n\5\3\6\3\6\3\6\3\6\3\6\3"+
		"\7\3\7\3\7\3\b\3\b\3\b\5\b>\n\b\3\b\3\b\3\b\3\t\3\t\6\tE\n\t\r\t\16\t"+
		"F\3\n\3\n\3\n\3\n\3\n\3\n\7\nO\n\n\f\n\16\nR\13\n\3\n\2\2\13\2\4\6\b\n"+
		"\f\16\20\22\2\2\2Q\2\24\3\2\2\2\4$\3\2\2\2\6*\3\2\2\2\b\60\3\2\2\2\n\62"+
		"\3\2\2\2\f\67\3\2\2\2\16:\3\2\2\2\20B\3\2\2\2\22H\3\2\2\2\24\25\7\24\2"+
		"\2\25\26\7\3\2\2\26\27\7\4\2\2\27\30\7\20\2\2\30\37\7\4\2\2\31\32\7\22"+
		"\2\2\32\33\7\4\2\2\33\34\7\20\2\2\34\36\7\4\2\2\35\31\3\2\2\2\36!\3\2"+
		"\2\2\37\35\3\2\2\2\37 \3\2\2\2 \"\3\2\2\2!\37\3\2\2\2\"#\7\5\2\2#\3\3"+
		"\2\2\2$(\7\23\2\2%&\7\6\2\2&\'\7\21\2\2\')\7\7\2\2(%\3\2\2\2()\3\2\2\2"+
		")\5\3\2\2\2*+\7\b\2\2+,\5\4\3\2,\7\3\2\2\2-\61\5\6\4\2.\61\5\2\2\2/\61"+
		"\5\4\3\2\60-\3\2\2\2\60.\3\2\2\2\60/\3\2\2\2\61\t\3\2\2\2\62\63\7\t\2"+
		"\2\63\64\7\4\2\2\64\65\7\20\2\2\65\66\7\4\2\2\66\13\3\2\2\2\678\7\n\2"+
		"\289\5\b\5\29\r\3\2\2\2:;\7\13\2\2;=\7\17\2\2<>\5\n\6\2=<\3\2\2\2=>\3"+
		"\2\2\2>?\3\2\2\2?@\7\n\2\2@A\5\b\5\2A\17\3\2\2\2BD\7\f\2\2CE\5\16\b\2"+
		"DC\3\2\2\2EF\3\2\2\2FD\3\2\2\2FG\3\2\2\2G\21\3\2\2\2HI\7\r\2\2IJ\7\16"+
		"\2\2JK\7\3\2\2KL\7\17\2\2LP\7\5\2\2MO\5\16\b\2NM\3\2\2\2OR\3\2\2\2PN\3"+
		"\2\2\2PQ\3\2\2\2Q\23\3\2\2\2RP\3\2\2\2\b\37(\60=FP";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}