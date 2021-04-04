import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;
import static org.junit.Assert.*;
import org.junit.Ignore;

/**
 *
 * @author ub
 */
public class TestRandomizedQueue {
    RandomizedQueue<Integer> randQueue;
    
    public TestRandomizedQueue() {
    }
    
    @BeforeClass
    public static void setUpClass() {
    }
    
    @AfterClass
    public static void tearDownClass() {
    }
    
    @Before
    public void setUp() {
        randQueue = new RandomizedQueue<>();
    }
    
    @After
    public void tearDown() {
    }
    
    @Test
    public void iter() {
        // to be or not to be
        randQueue.enqueue(1);
        randQueue.enqueue(2);
        randQueue.enqueue(3);
        randQueue.enqueue(4);
        randQueue.enqueue(5);
        
        for (Integer i: randQueue) {
            System.out.println(i);
        }
        
    }


    @Test
    @Ignore
    public void addLoad2() {
        int n = 10000000;
        for (int i = 0; i < n; i++) {
            randQueue.enqueue(i);
        }

        assert (randQueue.size() == n);
        
        for (int i = 0; i < n; i++) {
            randQueue.dequeue();
        }
        
        assert (randQueue.size() == 0);
    }
}
