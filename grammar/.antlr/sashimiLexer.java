// Generated from c:\gitrepos\sashimi\grammar\Sashimi.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SashimiLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, DIRECTIVE=11, COMMAND=12, IDENT=13, ALIAS=14, CONSTRAINT=15, 
		SEPERATOR=16, TYPE=17, UNION=18, WS=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"T__9", "DIRECTIVE", "COMMAND", "IDENT", "ALIAS", "CONSTRAINT", "SEPERATOR", 
		"TYPE", "UNION", "WS"
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


	public SashimiLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Sashimi.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\25\u00b2\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\5\r"+
		"y\n\r\3\16\6\16|\n\16\r\16\16\16}\3\17\6\17\u0081\n\17\r\17\16\17\u0082"+
		"\3\20\6\20\u0086\n\20\r\20\16\20\u0087\3\21\3\21\3\22\3\22\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u009d"+
		"\n\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u00aa"+
		"\n\23\3\24\6\24\u00ad\n\24\r\24\16\24\u00ae\3\24\3\24\2\2\25\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25\3\2\4\3\2c|\5\2\13\f\17\17\"\"\2\u00bc\2\3\3\2\2\2\2\5\3"+
		"\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2"+
		"\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3"+
		"\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'"+
		"\3\2\2\2\3)\3\2\2\2\5+\3\2\2\2\7-\3\2\2\2\t/\3\2\2\2\13\61\3\2\2\2\r\63"+
		"\3\2\2\2\17;\3\2\2\2\21>\3\2\2\2\23A\3\2\2\2\25C\3\2\2\2\27F\3\2\2\2\31"+
		"x\3\2\2\2\33{\3\2\2\2\35\u0080\3\2\2\2\37\u0085\3\2\2\2!\u0089\3\2\2\2"+
		"#\u009c\3\2\2\2%\u00a9\3\2\2\2\'\u00ac\3\2\2\2)*\7*\2\2*\4\3\2\2\2+,\7"+
		"$\2\2,\6\3\2\2\2-.\7+\2\2.\b\3\2\2\2/\60\7]\2\2\60\n\3\2\2\2\61\62\7_"+
		"\2\2\62\f\3\2\2\2\63\64\7n\2\2\64\65\7k\2\2\65\66\7u\2\2\66\67\7v\2\2"+
		"\678\7\"\2\289\7q\2\29:\7h\2\2:\16\3\2\2\2;<\7c\2\2<=\7u\2\2=\20\3\2\2"+
		"\2>?\7k\2\2?@\7u\2\2@\22\3\2\2\2AB\7/\2\2B\24\3\2\2\2CD\7q\2\2DE\7h\2"+
		"\2E\26\3\2\2\2FG\7u\2\2GH\7c\2\2HI\7u\2\2IJ\7j\2\2JK\7k\2\2KL\7o\2\2L"+
		"M\7k\2\2MN\7<\2\2N\30\3\2\2\2OP\7o\2\2PQ\7c\2\2QR\7t\2\2RS\7m\2\2ST\7"+
		"f\2\2TU\7q\2\2UV\7y\2\2Vy\7p\2\2WX\7f\2\2XY\7k\2\2YZ\7u\2\2Z[\7r\2\2["+
		"\\\7n\2\2\\]\7c\2\2]y\7{\2\2^_\7t\2\2_`\7g\2\2`a\7r\2\2ab\7g\2\2bc\7c"+
		"\2\2cy\7v\2\2de\7n\2\2ef\7c\2\2fg\7{\2\2gh\7q\2\2hi\7w\2\2ij\7v\2\2jk"+
		"\7a\2\2kl\7u\2\2lm\7g\2\2mn\7e\2\2no\7v\2\2op\7k\2\2pq\7q\2\2qy\7p\2\2"+
		"rs\7g\2\2st\7p\2\2tu\7v\2\2uv\7k\2\2vw\7v\2\2wy\7{\2\2xO\3\2\2\2xW\3\2"+
		"\2\2x^\3\2\2\2xd\3\2\2\2xr\3\2\2\2y\32\3\2\2\2z|\t\2\2\2{z\3\2\2\2|}\3"+
		"\2\2\2}{\3\2\2\2}~\3\2\2\2~\34\3\2\2\2\177\u0081\t\2\2\2\u0080\177\3\2"+
		"\2\2\u0081\u0082\3\2\2\2\u0082\u0080\3\2\2\2\u0082\u0083\3\2\2\2\u0083"+
		"\36\3\2\2\2\u0084\u0086\t\2\2\2\u0085\u0084\3\2\2\2\u0086\u0087\3\2\2"+
		"\2\u0087\u0085\3\2\2\2\u0087\u0088\3\2\2\2\u0088 \3\2\2\2\u0089\u008a"+
		"\7.\2\2\u008a\"\3\2\2\2\u008b\u008c\7v\2\2\u008c\u008d\7g\2\2\u008d\u008e"+
		"\7z\2\2\u008e\u009d\7v\2\2\u008f\u0090\7r\2\2\u0090\u0091\7k\2\2\u0091"+
		"\u0092\7e\2\2\u0092\u0093\7v\2\2\u0093\u0094\7w\2\2\u0094\u0095\7t\2\2"+
		"\u0095\u009d\7g\2\2\u0096\u0097\7p\2\2\u0097\u0098\7w\2\2\u0098\u0099"+
		"\7o\2\2\u0099\u009a\7d\2\2\u009a\u009b\7g\2\2\u009b\u009d\7t\2\2\u009c"+
		"\u008b\3\2\2\2\u009c\u008f\3\2\2\2\u009c\u0096\3\2\2\2\u009d$\3\2\2\2"+
		"\u009e\u009f\7o\2\2\u009f\u00a0\7c\2\2\u00a0\u00a1\7p\2\2\u00a1\u00a2"+
		"\7{\2\2\u00a2\u00a3\7Q\2\2\u00a3\u00aa\7h\2\2\u00a4\u00a5\7q\2\2\u00a5"+
		"\u00a6\7p\2\2\u00a6\u00a7\7g\2\2\u00a7\u00a8\7Q\2\2\u00a8\u00aa\7h\2\2"+
		"\u00a9\u009e\3\2\2\2\u00a9\u00a4\3\2\2\2\u00aa&\3\2\2\2\u00ab\u00ad\t"+
		"\3\2\2\u00ac\u00ab\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae\u00ac\3\2\2\2\u00ae"+
		"\u00af\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0\u00b1\b\24\2\2\u00b1(\3\2\2\2"+
		"\n\2x}\u0082\u0087\u009c\u00a9\u00ae\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}