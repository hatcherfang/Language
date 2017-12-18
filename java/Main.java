public class Main {

    public static void bubbleSort(int [] array){
        // eg:
        // array = [4, 5, 1, 3, 10, 6]
        int len;
        len = array.length;
        System.out.println("before sort:");
        for (int i = 0; i < len; i ++){
            System.out.print(array[i] + " ");
        }
        for (int i = 0; i < len; i ++){
            for (int j=0; j < len-i; j++) {
                if (j + 1 < len - i && array[j] > array[j + 1]) {
                    int tmp = array[j];
                    array[j] = array[j + 1];
                    array[j + 1] = tmp;
                }
            }
        }
        System.out.println("");
        System.out.println("after sort:");
        for (int i = 0; i < len; i ++){
           System.out.print(array[i] + " ");
        }
    }

    public static void insertSort(int[] array){
        if (array == null || array.length == 0){
            return;
        }
        int len = array.length;
        for (int i = 1; i < len; i ++){
            int tmp = array[i];
            int j = i;
            // move forward back
            for (; j > 0 && tmp < array[j-1]; j --){
                array[j] = array[j-1];
            }
            // insert
            array[j] = tmp;
        }
    }

    public static void qSort(int[] array, int left, int right){
        if (left >= right){
            return ;
        }
        int low = left;
        int high = right;
        int pivote = array[low];
        while (low < high){
            while (low < high && array[high] >= pivote){
                high --;
            }
            array[low] = array[high]; // move the little data to left
            while (low < high && array[low] < pivote){
                low ++;
            }
            array[high] = array[low]; // move the big data to right
        }
        array[low] = pivote; // assign the pivote to the middle
        qSort(array, left, low-1); // split left
        qSort(array, low+1, right); // split right
    }

    public static void quickSort(int[] array){
        if (array == null || array.length == 0){
            return;
        }
        qSort(array, 0, array.length-1);
    }

    public static void adjustHeap(int[] array, int len){
        int parent = len/2;
        while (parent > 0){
            int lchild = parent*2;
            int k = lchild;
            int rchild = parent*2 + 1;
            // first compare left and right child and get the bigger one
            if (rchild <= len && array[rchild-1] > array[lchild-1])
                k = rchild;
            // then compare the bigger child with parent node
            if (array[k-1] > array[parent-1]){
                int tmp = array[k-1];
                array[k-1] = array[parent-1];
                array[parent-1] = tmp;
            }
            parent --;
        }
        int tmp = array[0];
        array[0] = array[len-1];
        array[len-1] = tmp;
    }

    public static void heapSort(int[] array){
        if (array == null || array.length == 0){
            return;
        }
        int len = array.length;
        // each time get the biggest node and put it at the end
        for (int i = len; i > 0; i --){
            adjustHeap(array, i);
        }
    }

    public static void shellSort(int[] array){
        if (array == null || array.length == 0){
            return;
        }
        int len = array.length;
        for (int step = len/2; step > 0; step = step/2){
            // first set the step
            for (int i = step; i < len; i = i + step){
                int tmp = array[i];
                int j = i;
                // second sort the data according to the data serialize from back to front
                for (; j > 0 && tmp < array[j-step]; j = j - step){
                    array[j] = array[j-step];
                }
                if (j != i){
                    array[j] = tmp;
                }
            }
        }
    }

    public static void merge(int[] array, int left, int mid, int right){
        int tmp[] = new int[right-left+1];
        int i1 = left;
        int j1 = mid + 1;
        int k = 0;
        while (i1 <= mid && j1 <= right){
           if (array[i1] > array[j1]){
               tmp[k++] = array[j1++];
           }
           else {
               tmp[k++] = array[i1++];
           }
        }
        while (i1 <= mid){
            tmp[k++] = array[i1++];
        }
        while (j1 <= right){
            tmp[k++] = array[j1++];
        }
        k = 0;
        for (int i = left; i <= right; i ++, k ++){
            array[i] = tmp[k];
        }
    }

    public static void mergeSort(int[] array, int left, int right){
        if (array == null || array.length == 0){
            return;
        }
        if (left >= right)
            return;
        int mid = (left + right)/2;
        mergeSort(array, left, mid); // merge sort left part
        mergeSort(array, mid+1, right); // merge sort right part
        merge(array, left, mid, right); // merge the array from left to right
    }

    public static void main(String[] args) {
        // helloWorld();
        int NUM = 100;
        int array[] = new int[NUM];
        for (int i = 0; i < NUM; i ++){
            array[i] = (int)(Math.random()*NUM);
            System.out.print(array[i] + " ");
        }
        System.out.println("\ninsert sort:");
        insertSort(array);
        for (int i = 0; i < NUM; i ++){
            System.out.print(array[i] + " ");
        }
        System.out.println("\nquick sort:");
        quickSort(array);
        for (int i = 0; i < NUM; i ++){
            System.out.print(array[i] + " ");
        }
        System.out.println("\nheap sort:");
        heapSort(array);
        for (int i = 0; i < NUM; i ++){
            System.out.print(array[i] + " ");
        }
        System.out.println("\nshell sort:");
        shellSort(array);
        for (int i = 0; i < NUM; i ++){
            System.out.print(array[i] + " ");
        }
        System.out.println("\nmerge sort:");
        mergeSort(array, 0, array.length-1);
        for (int i = 0; i < NUM; i ++){
            System.out.print(array[i] + " ");
        }
        System.out.println("\nbubble sort:");
        bubbleSort(array);
        for (int i = 0; i < NUM; i ++){
            System.out.print(array[i] + " ");
        }
    }
}
