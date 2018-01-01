
import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Ignore;
import org.junit.Test;

public class TestDeque {

    Deque<String> stringQueue;
    Deque<Integer> intQueue;

    public TestDeque() {
    }

    @BeforeClass
    public static void setUpClass() {
    }

    @AfterClass
    public static void tearDownClass() {
    }

    @Before
    public void setUp() {
        stringQueue = new Deque<>();
        intQueue = new Deque<>();

    }

    @After
    public void tearDown() {
    }

    @Test
    public void shoudlBeEmptry() {
        assert (stringQueue.isEmpty());
    }

    @Test
    public void addRemove() {
        // to be or not to be
        stringQueue.addFirst("or");
        stringQueue.addFirst("be");
        stringQueue.addFirst("to");
        stringQueue.addLast("not");
        stringQueue.addLast("to");
        stringQueue.addLast("be");

        assert (stringQueue.size() == 6);

        // remove all items        
        int k = stringQueue.size();
        for (int i = 0; i < k; i++) {
            stringQueue.removeLast();
        }
        assert (stringQueue.isEmpty());
        
        stringQueue.addFirst("one");
        assert (stringQueue.size() == 1);

        String item =  stringQueue.removeLast();
        
        assert (item.equals("one"));
        
        stringQueue.addFirst("two");
        assert (stringQueue.size() == 1);
    }
    
    @Test
    public void iter() {
        // to be or not to be
        stringQueue.addFirst("or");
        stringQueue.addFirst("be");
        stringQueue.addFirst("to");
        stringQueue.addLast("not");
        stringQueue.addLast("to");
        stringQueue.addLast("be");
        
        for (String s: stringQueue) {
            System.out.println(s);
        }
        for (String s: stringQueue) {
            System.out.println(s);
        }
        
    }

    @Test
    @Ignore
    public void addLoad() {
        int n = 10000000;
        for (int i = 0; i < n; i++) {
            intQueue.addFirst(i);
        }

        assert (intQueue.size() == n);
        //System.out.println(ObjectSizeFetcher.getObjectSize(intQueue));

//        for (int i = 0; i < n; i++) {
//            intQueue.removeLast();
//        }
    }

    @Test
    public void memTest() {
        //java.lang.instrument.Instrumentation instrumentation;

    }
}
