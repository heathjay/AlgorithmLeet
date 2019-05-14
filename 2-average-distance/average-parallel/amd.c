#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>
#include <limits.h>

//Linux
#include <pthread.h>
#include <assert.h>
#include <sys/sysinfo.h>

pthread_barrier_t barrier;


struct thread_work{
   int id;
   int start;
   int end;
   int v;
   uint32_t* dists;
};



//14.05
void* md_all_pairs(void *tw) {
    struct thread_work *thread_work_d = (struct thread_work*)tw;
    int id = thread_work_d  -> id;
    int start = thread_work_d  -> start;
    int end = thread_work_d  -> end;
    int v = thread_work_d  ->v;
    uint32_t* dists = thread_work_d  -> dists;
    uint32_t k, i, j;

   
    for (k = 0; k < v; ++k) {
        
       int flag = pthread_barrier_wait(&barrier);
 		if(flag != 0 && flag != PTHREAD_BARRIER_SERIAL_THREAD)
 		{
 			printf("Could not wait on barrier\n");
 			exit(-1);
 		}
         printf("thread %d works\n",id);
        for (i = start; i < end; ++i) {
            for (j = 0; j < v; ++j) {
                uint32_t intermediary = dists[i*v+k] + dists[k*v+j];
                // Checks for overflows
                if ((intermediary >= dists[i*v+k])&&(intermediary >= dists[k*v+j])&&(intermediary < dists[i*v+j]))
                    dists[i*v+j] = dists[i*v+k] + dists[k*v+j];
            }
        }
    }

}




/* Computes the average minimum distance between all pairs of vertices with a path connecting them */
void amd (uint32_t* dists, uint32_t v) {
    uint32_t i, j;
	//uint32_t infinity = v*v;
    uint32_t infinity = 1000;
	uint32_t smd = 0; 	//sum of minimum distances
	uint32_t paths = 0; //number of paths found
	uint32_t solution = 0;

    for (i = 0; i < v; ++i) {
        for (j = 0; j < v; ++j) {
			// We only consider if the vertices are different and there is a path
            if ((i != j) && (dists[i*v+j] < infinity)) {
				smd += dists[i*v+j];
				paths++;
			}
        }
    }

	solution = smd / paths;
	printf("%d\n", solution);

}



/* Debug function (not to be used when measuring performance)*/
void debug (uint32_t* dists, uint32_t v) {
    uint32_t i, j;
	// uint32_t infinity = v*v;
    uint32_t infinity = 1000;
    for (i = 0; i < v; ++i) {
        for (j = 0; j < v; ++j) {
            if (dists[i*v+j] > infinity) printf("%7s", "inf");
            else printf ("%7u", dists[i*v+j]);
        }
        printf("\n");
    }
}

// Main program - reads input, calls FW, shows output
int main (int argc, char* argv[]) {
  
    //Reads input 
    //First line: v (number of vertices) and e (number of edges)
    uint32_t v, e;
    printf("input v and e\n");
    scanf("%u %u", &v, &e);  

	//13.05
	printf("v:%u e:%u\n",v,e);

    //14.05 add cpus
    int cpus = get_nprocs();
    if (getenv("MAX_CPUS")){
        cpus = atoi(getenv("MAX_CPUS"));
    }
    assert(cpus > 0 && cpus <= 64);
    printf("Running on the %d cpus\n",cpus);
    
    //thread = (pthread_t*)malloc(cpus * sizeof(pthread_t));
    pthread_t thread[cpus];
    struct thread_work thread_work_d[cpus];
    if(pthread_barrier_init(&barrier,NULL, cpus)){
        printf("barrier initialization error\n");
        return -1;
    }

    //Allocates distances matrix (w/ size v*v) i
    //and sets it with max distance and 0 for own vertex
    uint32_t* dists = malloc(v*v*sizeof(uint32_t));
    memset(dists, UINT32_MAX, v*v*sizeof(uint32_t));
    uint32_t i;
    //13.05
    for ( i = 0; i < v; ++i ) dists[i*v+i] = 0;
    //13.05

    //Reads edges from file and sets them in the distance matrix
    uint32_t source, dest, cost;
    for ( i = 0; i < e; ++i ){
        printf("input source dest cost\n");
        scanf("%u %u %u", &source, &dest, &cost);
        if (cost < dists[source*v+dest]) dists[source*v+dest] = cost;  
        //13.05
        printf("source:%u dest:%u cost:%u\n",source ,dest,cost);

    }
    //14.05
    for(int i=0; i < cpus; i++){
         thread_work_d[i].id = i;
         thread_work_d[i].v = v;
         thread_work_d[i].start = i * v / cpus;
         thread_work_d[i].dists = dists;
        if (i == cpus -1){     
            thread_work_d[i].end = v;    
        }else{
            thread_work_d[i].end = (i + 1) * v / cpus;
        }

        if(pthread_create(&thread[i],NULL,md_all_pairs,(void*)&thread_work_d[i])){
            printf("create %d thread error\n",i);
        }
    }

    for (i = 0; i < cpus  ; i++){
 		pthread_join (thread[i], NULL);
    }

	//Computes the minimum distance for all pairs of vertices
    //md_all_pairs(dists, v);
    //13.05
    //for ( i = 0; i < v*v; ++i ) printf("%d",dists[i*v+i]);
    //Computes and prints the final solution
    amd(dists, v);

#if DEBUG
	debug(dists, v);
#endif

    return 0;
}
